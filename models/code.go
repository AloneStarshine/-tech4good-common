package models

//ErrorValue 错误值配置
type ErrorValue struct {
	Code int
	Msg  string
	Log  string
}

//ErrorCode 全局返回码
var ErrorCode = map[string]ErrorValue{
	"success":         {Code: 0, Msg: "", Log: ""},
	"SerError":        {Code: 100001, Msg: "服务请求失败！", Log: "服务【%s】请求失败：%s"},
	"SerResEmpty":     {Code: 100002, Msg: "服务请求返回为空！", Log: "服务【%s】请求返回为空！"},
	"SerResJosnError": {Code: 100003, Msg: "服务请求返回内容错误！", Log: "服务【%s】请求返回内容错误：%s"},
}
