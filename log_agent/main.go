// 日志采集客户端
// https://blog.csdn.net/Bobdragery/article/details/106093951
package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main(){
	fmt.Println(os.Getpid())
	lancer := New()
	lancer.Process()
	c := make(chan os.Signal, 1)
	signal.Notify(c,syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for{
		s := <-c
		fmt.Printf("get a signal %s\n", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			lancer.Close()
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}


type LancerConfig struct {
	SocketPath string `json:"socket_path"` // unix domain socket 管道路径
}

type LancerAgent struct {
	Cfg *LancerConfig
	Listener net.Listener
	StopChan chan struct{}
}

// 创建实例
func New() *LancerAgent{
	logConfig := &LancerConfig{
		SocketPath:"/Users/shenxinhua/workspace/data/log.sock",
	}
	lancer := &LancerAgent{
		Cfg:logConfig,
		StopChan: make(chan struct{},1),
	}
	// 清理sock文件
	os.Remove(lancer.Cfg.SocketPath)
	// 监听sock
	var err error
	lancer.Listener, err= net.Listen("unix",lancer.Cfg.SocketPath)
	if err != nil {
		fmt.Println("error1:",err)
		panic(err)
	}
	fmt.Println("lancer is listened ",lancer.Cfg.SocketPath)
	return lancer
}
// 处理日志
func (lancer *LancerAgent) Process(){
	f := func(lancer *LancerAgent){
		// 接收连接
		for {
			conn,err := lancer.Listener.Accept()
			if err != nil {
				fmt.Println("error2:",err)
				// 优雅退出
				select {
				case <-lancer.StopChan:
					return
				default:
				}
				return
			}
			// 处理日志，这里需要考虑日志量大如何处理，todo 后续建议走协程池
			go func(c net.Conn) {
				defer c.Close()
				reader := bufio.NewReader(c)
				msg,err := Decode(reader)
				if err != nil{
					fmt.Println("error:",err)
				}else{
					fmt.Println("recv: ",msg)
				}
				/*buf := make([]byte,16)
				for{
					// todo 粘包问题
					_, err := c.Read(buf)
					if err != nil {
						if err != io.EOF {
							c.Close()
							break
						}else{
							fmt.Println("end...")
							break
						}
					}
				}
				fmt.Println("recv: ", string(buf))*/
			}(conn)
		}
	}
	go f(lancer)
}

// Encode 消息编码协议
/*func Encode(message string) ([]byte,error){
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
}*/

// Decode 解码消息
func Decode(reader *bufio.Reader)(string,error){
	// 读取4字节的数据，得到包的内容长度
	lengthByte,_ := reader.Peek(4)
	// 定义一个以lengthByte为内容的缓冲区
	lengthBuffer := bytes.NewBuffer(lengthByte)
	var length int32
	err := binary.Read(lengthBuffer,binary.LittleEndian,&length)
	if err != nil {
		return "",err
	}
	// Buffered返回缓冲区现有的可读取的字节数，前面用Peek读取，所以这里数据内容应该大于 length+4
	if int32(reader.Buffered()) < length + 4 {
		return "",err
	}
	// 读取真正的消息数据
	pack := make([]byte,int(4+length))
	_,err = reader.Read(pack)
	if err != nil {
		return "",err
	}
	return string(pack[4:]),nil
}

// 优雅关停
func (lancer *LancerAgent) Close(){
	lancer.StopChan <- struct{}{}
	lancer.Listener.Close()
	time.Sleep(time.Second)
	fmt.Printf("lancer is closed.")
}