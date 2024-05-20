package app

import (
	"fmt"
	"net/http"
)

type Error struct {
	code   int      // 错误码
	msg    string   // 错误消息
	detail []string // 详细信息
}

func NewError(code int, msg string) *Error {
	return &Error{code: code, msg: msg}
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

func (e *Error) Details() []string {
	return e.detail
}

func (e *Error) WithDetails(details ...string) *Error {
	ne := *e
	ne.detail = make([]string, len(details))
	for i, d := range details {
		ne.detail[i] = d
	}

	return &ne
}

func (e *Error) StatusCode() int {
	switch e.Code() {
	case SuccessOk.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code(), UnauthorizedTokenError.Code(),
		UnauthorizedTokenGenerate.Code(), UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests

	default:
		return http.StatusInternalServerError
	}
}
