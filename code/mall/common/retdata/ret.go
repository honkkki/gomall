package retdata

import "github.com/honkkki/gomall/code/mall/common/errorx"

type SuccessRet struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewSuccessRet(data interface{}) SuccessRet {
	return SuccessRet{
		Code: errorx.Success,
		Msg:  "success",
		Data: data,
	}
}
