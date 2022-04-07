package context

import (
	"git.code.oa.com/fip-team/rasse/config"
	"git.code.oa.com/fip-team/rasse/xgin"
	"git.code.oa.com/fip-team/rasse/xlog"
	"github.com/gin-gonic/gin"
)

// BaseContext BaseContext
type BaseContext struct {
	TraceID      string
	Tech4GoodLog xlog.ILogger
	HttpCtx      *gin.Context
}

// WithLogCatalog WithLogCatalog
func (ctx *BaseContext) WithLogCatalog(catalog string) xlog.ILogger {
	var defaultField = make(map[string]interface{})
	defaultField["trace_id"] = ctx.TraceID
	if ctx.HttpCtx != nil {
		defaultField["path"] = ctx.HttpCtx.Request.URL.Path
	}
	log := xlog.WithCataLog(catalog).WithFields(defaultField)
	return log
}

// InitContext 初始化上下文
func (ctx *BaseContext) InitContext(context *gin.Context) error {
	xgin.InitTraceID(context)
	ctx.TraceID = xgin.GetTraceID(context)
	log := ctx.WithLogCatalog(config.GetStringOrDefault("log.catalog", xlog.DefaultCataLog))
	ctx.Tech4GoodLog = log
	ctx.HttpCtx = context
	return nil
}
