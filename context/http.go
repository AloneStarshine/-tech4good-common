package context

import (
	"git.code.oa.com/fip-team/rasse/xgin"
	"git.code.oa.com/fip-team/rasse/xhttp"

	"git.code.oa.com/fip-team/rasse/extension/xstring"
	"git.code.oa.com/tech4good/common/constant"
)

// InitRequest 初始化请求
func (ctx *BaseContext) InitRequest() *xhttp.HTTPAgent {
	svc := xhttp.New().
		SetHeader(xgin.CtxKey_TraceID, ctx.TraceID).
		SetHeader(constant.NONCE, xstring.GetUUID())
	return svc
}
