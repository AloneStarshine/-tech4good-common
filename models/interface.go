package models

import "git.code.oa.com/tech4good/common/error"

// 这里那么多new方法是类似于java的类构造函数
// 由于go不支持重载，不支持struct构造函数，会用方法名进行区分

// Response 接口输出类
type Response struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data,omitempty"`
	ErrMsg interface{} `json:"errmsg,omitempty"`
}

//以下为创建并返回标准格式不同种类的response实例，方便直接填入内容并发送

// NewDefaultResponse 默认 Response
func NewDefaultResponse() *Response {
	return &Response{
		Code:   0,
		Data:   "",
		ErrMsg: "",
	}
}

// NewErrorResponse 错误 Response
func NewErrorResponse(s *error.Status) *Response {
	return &Response{
		Code:   int(s.Code),
		Data:   "",
		ErrMsg: s.Message,
	}
}

// NewResponse 自定义 Response
func NewResponse(code int, data interface{}, err interface{}) *Response {
	return &Response{
		Code:   code,
		Data:   data,
		ErrMsg: err,
	}
}
