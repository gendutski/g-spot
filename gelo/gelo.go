package gelo

import "strings"

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func Init(code int, messages ...string) *Error {
	res := &Error{code, ""}
	if len(messages) > 0 {
		res.Message = strings.Join(messages, "; ")
	}
	return res
}
