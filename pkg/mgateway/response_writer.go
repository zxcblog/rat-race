package mgateway

import (
	"fmt"
	"net/http"
)

type responseWriter struct {
	http.ResponseWriter
	status int
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Error struct {
	code int    `json:"code"`
	msg  string `json:"msg"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}
