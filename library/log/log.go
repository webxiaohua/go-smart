package log

import (
	"context"
	"sync"
)

type Level int
const (
	_debugLevel Level = iota
	_infoLevel
	_warnLevel
	_errorLevel
	_fatalLevel
)
func (l Level) String() string {
	return levelNames[l]
}

type Handler interface {
	Log(context.Context, Level)
	SetFormat(string)
	Close() error
}

var (
	levelNames = [...]string{
		_debugLevel: "DEBUG",
		_infoLevel:  "INFO",
		_warnLevel:  "WARN",
		_errorLevel: "ERROR",
		_fatalLevel: "FATAL",
	}
	_lock sync.RWMutex
	_handler Handler
)

func Init() {

}

func h() (handler Handler) {
	_lock.RLock()
	handler = _handler
	_lock.RUnlock()
	return
}

