package errorx

import "net/http"

const defaultCode = RequestError

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewCodeError(code int, msg string) error {
	return CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e CodeError) Error() string {
	return e.Msg
}

func (e CodeError) Response() CodeErrorResponse {
	return CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func ErrorHandler() func(err error) (int, interface{}) {
	return func(err error) (int, interface{}) {
		switch e := err.(type) {
		case CodeError:
			return http.StatusOK, e.Response()
		default:
			exi := NewCodeError(InternalError, e.Error())
			ex, ok := exi.(CodeError)
			if !ok {
				return http.StatusInternalServerError, e.Error()
			}
			return http.StatusOK, ex.Response()
		}
	}
}
