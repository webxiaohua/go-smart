package dsn

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T){
	resp,err := Parse("https://www.baidu.com/search?a=1")
	if err != nil {
		fmt.Println("[error]",err)
	}else{
		fmt.Printf("schema:%s host:%s port:%s \n",resp.Scheme,resp.Host,resp.Port())
	}
}