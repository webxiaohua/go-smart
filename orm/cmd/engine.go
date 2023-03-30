// https://mp.weixin.qq.com/s/fk__wCp-r60FJwYPkZpmFQ?utm_source=wechat_session&utm_medium=social&utm_oi=27739344076800

package main

import (
	"database/sql"
	"errors"

	//"go-smart/library/errors"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

var (
	//ERR_BATCH_INSERT_PARAMTYPE = errors.New(400,"批量插入的子元素必须是结构体类型")
)

type SmartOrmEngine struct {
	DB *sql.DB
	TableName string
	Prepare string
	AllExec []interface{}
	Sql string
	WhereParam   string
	LimitParam   string
	OrderParam   string
	OrWhereParam string
	WhereExec    []interface{}
	UpdateParam  string
	UpdateExec   []interface{}
	FieldParam   string
	TransStatus  int
	Tx           *sql.Tx
	GroupParam   string
	HavingParam  string
}

func NewMySql(Username string, Password string, Address string, Dbname string) (*SmartOrmEngine, error) {
	dsn := Username + ":" + Password + "@tcp(" + Address + ")/" + Dbname + "?charset=utf8&timeout=5s&readTimeout=6s"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//最大连接数等配置，先占个位
	//db.SetMaxOpenConns(3)
	//db.SetMaxIdleConns(3)
	return &SmartOrmEngine{
		DB:         db,
		FieldParam: "*",
	}, nil
}
// 条件
func (s *SmartOrmEngine) Where(name string) *SmartOrmEngine {
	return s
}
// 限制条件
func (s *SmartOrmEngine) Limit(name string) *SmartOrmEngine {
	return s
}
// 指定表
func (s *SmartOrmEngine) Table(name string) *SmartOrmEngine {
	s.TableName = name
	s.resetSmartOrmEngine()
	return s
}
// 重置引擎
func (s *SmartOrmEngine) resetSmartOrmEngine(){

}

// 获取表名
func (s *SmartOrmEngine) GetTable(name string) string {
	return s.TableName
}

// 新增数据，支持单个和批量
func (s *SmartOrmEngine) Insert(data interface{}) (lastId int64,err error){
	reflectData := reflect.ValueOf(data)
	if reflectData.Kind() == reflect.Slice {
		// 批量
		return s.batchInsert(data)
	}else{
		// 单个
		return s.singleInsert(data)
	}
}

// 批量新增
func (s *SmartOrmEngine) batchInsert(batchData interface{}) (lastId int64,err error) {
	// 反射解析
	getValue := reflect.ValueOf(batchData)
	// 切片大小
	l := getValue.Len()
	// 字段名
	var fieldName []string
	// 占位符
	var placeholderString []string
	// 循环判断
	for i:=0;i<l;i++{
		value := getValue.Index(i) // value of item
		typed := value.Type() // type of item
		if typed.Kind() != reflect.Struct {
			err = errors.New("批量插入的子元素必须是结构体类型")
			return
		}
		num := value.NumField()
		// 子元素值
		var placeholder []string
		// 遍历子元素
		for j:=0;j<num;j++ {
			// 小写开头 无法反射，跳过
			if !value.Field(j).CanInterface() {
				continue
			}
			// 解析tag，找到真实的field
			sqlTag := typed.Field(j).Tag.Get("sql")
			if sqlTag != "" {
				// 跳过自增字段
				if strings.Contains(sqlTag,"auto_increment") {
					continue
				}
				// 字段名只取第一个
				if i == 0 {
					fieldName = append(fieldName,strings.Split(sqlTag,",")[0])
				}
				placeholder = append(placeholder,"?")
			}else{
				// 字段名只取第一个
				if i == 0 {
					fieldName = append(fieldName,typed.Field(j).Name)
				}
				placeholder = append(placeholder,"?")
			}
			// 字段值
			s.AllExec = append(s.AllExec,value.Field(j).Interface())
		}
		//子元素拼接成多个()括号后的值
		placeholderString = append(placeholderString,"("+strings.Join(placeholder,",")+")")
	}
	// 拼接表、字段名、占位符
	s.Prepare = "INSERT INTO "+s.TableName+" ("+strings.Join(fieldName,",")+") VALUES " + strings.Join(placeholderString,",")
	// prepare
	var stmt *sql.Stmt
	stmt,err = s.DB.Prepare(s.Prepare)
	// 执行exec
	result,err := stmt.Exec(s.AllExec...)
	if err != nil {
		return 0,s.setErrorInfo(err)
	}
	lastId,err = result.LastInsertId()
	return
}
// 单个新增
func (s *SmartOrmEngine) singleInsert(data interface{})(lastId int64,err error){
	var fields []string
	var placeholders []string
	var params []interface{}
	reflectData := reflect.ValueOf(data)
	num := reflectData.NumField()
	for i:=0;i<num;i++ {

	}
}

func (s *SmartOrmEngine) setErrorInfo(err error) error {
	_, file, line, _ := runtime.Caller(1)
	return errors.New("File: " + file + ":" + strconv.Itoa(line) + ", " + err.Error())
}