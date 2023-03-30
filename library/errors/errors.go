

package errors

import "encoding/json"

type Err struct {
	Code int64
	Msg   string
}
func (e *Err) Error() string {
	err, _ := json.Marshal(e)
	return string(err)
}
func New(code int64, msg string) *Err {
	return &Err{
		Code: code,
		Msg:   msg,
	}
}