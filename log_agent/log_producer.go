package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Log struct {
	AppId string `json:"app_id"`
	Timestamp int64 `json:"timestamp"`
	Env string `json:"env"`
	EnvColor string `json:"env_color"`
	Cost int64 `json:"cost"` // 耗时毫秒数
	InstanceId string `json:"instance_id"`
	Level string `json:"level"`
	LevelValue int64 `json:"level_value"`
	Log string `json:"log"`
	Source string `json:"source"`
	Zone string `json:"zone"`
}

func main(){
	sockP := "/Users/shenxinhua/workspace/data/log.sock"
	conn,err := net.Dial("unix",sockP)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	logItem := &Log{
		AppId:      "你好中国人",
		Timestamp:  time.Now().Unix(),
		Env:        "uat",
		InstanceId: "test1-234298",
		Level:      "INFO",
		LevelValue: 0,
		Log:        "{\"name\":\"12js\"}",
		Source:     "/go/src/go-live/app/job/test01/internal/service/reward/comsumer.go:63",
		Zone:       "sh001",
	}
	logItemByte,_ := json.Marshal(logItem)
	encodeLogItemByte,_ := Encode(string(logItemByte))
	if _, err := conn.Write(encodeLogItemByte); err != nil {
		fmt.Println(err)
		return
	}else{
		fmt.Println("succeed")
	}
}

// Encode 消息编码协议
func Encode(message string) ([]byte,error){
	var length = int32(len(message)) // 读取消息长度，转换成int32类型（4字节）
	var pkg = new(bytes.Buffer) // 定义一个空bytes缓冲区
	// 写入消息头 通过小端序方式把length写到pkg
	err := binary.Write(pkg,binary.LittleEndian,length)
	if err != nil {
		return nil,err
	}
	// 写入消息体 通过小端序方式把内容包写入pkg，封装成一个包
	err = binary.Write(pkg,binary.LittleEndian,[]byte(message))
	if err != nil {
		return nil,err
	}
	return pkg.Bytes(),nil
}