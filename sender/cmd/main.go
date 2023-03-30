package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"path"
)

// 设置websocket
var upGrader = websocket.Upgrader{
	// CheckOrigin防止跨站点的请求伪造
	CheckOrigin:func(r *http.Request) bool{
		return true
	},
}

// websocket 功能实现
func ping(c *gin.Context){
	//升级get请求为webSocket协议
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer ws.Close() //返回前关闭
	for {
		//读取ws中的数据
		mt, message, err := ws.ReadMessage()
		if err != nil {
			break
		}
		//写入ws数据
		err = ws.WriteMessage(mt, message)
		if err != nil {
			break
		}
	}
}

func upload(c *gin.Context){
	file,err := c.FormFile("file")
	if err != nil {
		fmt.Println("error1:",err)
		return
	}
	dst := path.Join("./static",file.Filename)
	saveErr := c.SaveUploadedFile(file,dst)
	if saveErr != nil {
		fmt.Println("error2:",saveErr)
		return
	}else{
		c.JSON(http.StatusOK,gin.H{
			"code":0,
			"msg":"success",
			"data":dst,
		})
	}
}

func main(){
	r := gin.Default()
	r.GET("/ping",ping)
	r.Run(":8000")
}