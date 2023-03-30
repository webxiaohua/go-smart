package dsn

import (
	"github.com/go-playground/validator/v10"
	"net/url"
	"os"
	"strings"
)

// 检验struct数据
var _validator *validator.Validate

func init() {
	_validator = validator.New()
}

// DSN 结构内容复刻 url
type DSN struct {
	*url.URL
}

func (d *DSN) Addresses() []string {
	switch d.Scheme {
	case "unix", "unixgram", "unixpacket":
		return []string{d.Path}
	}
	return strings.Split(d.Host, ",")
}

// 解析dsn
func Parse(strDsn string) (*DSN, error) {
	strDsn = os.ExpandEnv(strDsn)
	u, err := url.Parse(strDsn)
	return &DSN{URL: u}, err
}