package error

import (
	"fmt"
	"strconv"
)

// 设计来源于gRPC中对error的自定义处理
// 只是简单复现，并没有实现原型模式
// 如果后续迭代有需求可以参考grpc源码

// Error 自定义error
type Error struct {
	Status *Status
}

// 格式化 Error 信息
func (e Error) Error() string {
	return fmt.Sprintf("error: code = %d desc = %s", e.Status.GetCode(), e.Status.GetMessage())
}

// GetCode 获取 error 对应的 Code，如自定义错误map中没有定义则默认返回 -1 OtherError
func GetCode(err error) int {
	if se, ok := FromError(err); ok {
		c, err := strconv.Atoi(se.Code.String())
		if err != nil {
			return -1
		}
		return c
	}
	return -1
}

// GetMessage 获取 error 对应的 Message
func GetMessage(err error) string {
	if se, ok := FromError(err); ok {
		return se.Message
	}
	return ErrorMap[Unknown]
}
