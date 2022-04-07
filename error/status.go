package error

import (
	"fmt"
)

// Status 自定义错误模型，适用于不同环境
type Status struct {
	Code    Code   `json:"Code"`   // 错误码
	Message string `json:"ErrMsg"` // 错误信息
}

// Err 将 Status 转化为 Error
func (s *Status) Err() error {
	if s.Code == Success {
		return nil
	}
	return &Error{Status: s}
}

// GetCode 获取 Status 对应的 Code
func (s *Status) GetCode() Code {
	return s.Code
}

// GetMessage 获取 Status 对应的 Message
func (s *Status) GetMessage() string {
	return s.Message
}

// New 根据 Code 生成 Status ，可自定义错误信息 Message
//  [Example]
//  s := error.New(error.OtherError, "my error message")
//  fmt.Print(s.GetMessage())  // 输出："my error message"
//  fmt.Print(s.GetCode())  // 输出：-2
func New(code Code, msg ...string) *Status {
	return &Status{
		Code:    code,
		Message: code.GetMsg(msg...),
	}
}

// Newf 根据 Code 生成 Status ，可自定义format格式的错误信息 Message
//  [Example]
//  s := error.Newf(error.LoginTimeoutError, "%s login fail!", "XiaoMing")
//  fmt.Print(s.GetMessage())  // 输出："XiaoMing login fail!"
//  fmt.Print(s.GetCode())  // 输出：200006
func Newf(code Code, format string, a ...interface{}) *Status {
	return New(code, fmt.Sprintf(format, a...))
}

// FromError 将 error 转化为 Status
//  [Example]
//  _, err := strconv.Atoi("hello")
//	s, ok := FromError(err)
//	if ok {
//		fmt.Println(s.Code)  // 输出：-2
//		fmt.Println(s.Message) // 输出：strconv.Atoi: parsing "hello": invalid syntax
//	}
func FromError(err error) (s *Status, ok bool) {
	if err == nil {
		return nil, false
	}
	if se, ok := err.(*Error); ok {
		return se.Status, false
	}
	return New(OtherError, err.Error()), true
}
