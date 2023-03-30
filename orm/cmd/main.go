package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main(){
	/*db, err := sql.Open("mysql", "root:237k6pc9@tcp(175.24.187.251:3306)/test?charset=utf8&parseTime=true") //第一个参数数驱动名
	if err != nil {
		fmt.Println("error:",err)
	}else{
		fmt.Println("success")
	}
	list,err := queryList(db)
	if err != nil {
		fmt.Println("error2:",err)
	}
	fmt.Println(jsonString(list))
	db.Close()
	fmt.Printf("end")*/
	orm,err := NewMySql("root","237k6pc9","175.24.187.251:3306","test")
	if err != nil {
		fmt.Println("NewMySql error:",err)
		return
	}
	batchData := make([]TblMsg,0)
	batchData = append(batchData,TblMsg{
		Id:      0,
		Uid:     110000123,
		Content: "hello",
		Ctime:   time.Now(),
		Mtime:   time.Now(),
	})
	batchData = append(batchData,TblMsg{
		Id:      0,
		Uid:     110000124,
		Content: "hello2",
		Ctime:   time.Now(),
		Mtime:   time.Now(),
	})
	lastId,err := orm.Table("tbl_msg").BatchInsert(batchData)
	if err != nil {
		fmt.Println("BatchInsert error:",err)
		return
	}else{
		fmt.Println("lastId: ",lastId)
	}
}

func batchInsert(db *sql.DB){
	//rand := rand.New(rand.NewSource(time.Now().Unix()))
	for i:=0;i<5000;i++ {
		//randNum := rand.Int63n(10000000)
		//uid := 110000111+randNum
		content := fmt.Sprintf("hello %d",i)
		insertRows(db,&TblMsg{
			Uid:    110000111,
			Content: content,
		})
	}
}

// 消息表
type TblMsg struct {
	Id int64 `sql:"id,auto_increment"` //id
	Uid int64 `sql:"uid"` //用户id
	Content string `sql:"content"` //内容
	Ctime time.Time `sql:"ctime"` //创建时间
	Mtime time.Time `sql:"mtime"` //更新时间
}

func insertRows(db *sql.DB,msg *TblMsg) (res bool, err error){
	strSql := "INSERT INTO tbl_msg (`uid`, `content`) VALUES (?, ?)"
	var tmp sql.Result
	tmp,err = db.Exec(strSql,msg.Uid,msg.Content)
	if err == nil {
		rows,_ := tmp.RowsAffected()
		res = rows > 0
	}
	return
}

func insertLastId(db *sql.DB,msg *TblMsg)(lastId int64,err error){
	strSql := "INSERT INTO tbl_msg (`uid`, `content`) VALUES (?, ?)"
	stmt,stmtErr := db.Prepare(strSql)
	if stmtErr != nil {
		err = stmtErr
		return
	}
	dbRes,dbErr := stmt.Exec(msg.Uid,msg.Content)
	if dbErr != nil {
		err = dbErr
		return
	}
	lastId,err = dbRes.LastInsertId()
	return
}

func deleteRows(db *sql.DB,msg *TblMsg)(res bool,err error){
	strSql := "delete from tbl_msg where id = ?"
	var tmp sql.Result
	tmp,err = db.Exec(strSql,msg.Id)
	if err == nil {
		rows,_ := tmp.RowsAffected()
		res = rows > 0
	}
	return
}

func updateRows(db *sql.DB, msg *TblMsg)(res bool,err error){
	strSql := "update tbl_msg set uid=?,content=? where id=?"
	var tmp sql.Result
	tmp,err = db.Exec(strSql,msg.Uid,msg.Content,msg.Id)
	if err == nil {
		rows,_ := tmp.RowsAffected()
		res = rows > 0
	}
	return
}

func queryByPK(db *sql.DB,msg *TblMsg)(res *TblMsg,err error){
	res = &TblMsg{}
	strSql := "select id,uid,content,ctime,mtime from tbl_msg where id = ?"
	row := db.QueryRow(strSql,msg.Id)
	err = row.Scan(&res.Id,&res.Uid,&res.Content,&res.Ctime,&res.Mtime)
	return
}

func queryList(db *sql.DB)(res []*TblMsg,err error){
	res = make([]*TblMsg,0)
	strSql := "select id,uid,content,ctime,mtime from tbl_msg limit 10"
	var rows *sql.Rows
	rows,err = db.Query(strSql)
	if err == nil {
		defer rows.Close()
		tmpMsg := &TblMsg{}
		for rows.Next() {
			if err = rows.Scan(&tmpMsg.Id,&tmpMsg.Uid,&tmpMsg.Content,&tmpMsg.Ctime,&tmpMsg.Mtime); err != nil {
				return
			}
			res = append(res,tmpMsg)
		}
	}
	return
}

func jsonString(d interface{}) string {
	jsonByte,_ := json.Marshal(d)
	return string(jsonByte)
}
