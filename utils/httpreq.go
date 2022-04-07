package utils

import (
	"strings"

	"git.code.oa.com/fip-team/rasse/extension/xstring"
	"git.code.oa.com/fip-team/rasse/xcontext"
	"git.code.oa.com/fip-team/rasse/xgin"
	"git.code.oa.com/fip-team/rasse/xhttp"
	"git.code.oa.com/fip-team/rasse/xlog"
	"git.code.oa.com/tech4good/common/context"
	"git.code.oa.com/tech4good/common/models"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

// Http 带上下文的请求
func Http(svr string, url string, method string, expQuery string, timeOut int, headers map[string]string,
	body string, ctx *context.BaseContext, fill bool, errDetail bool) (*xhttp.Res, error, *models.Response) {
	var hlog xlog.ILogger
	var httpAgent *xhttp.HTTPAgent
	if ctx == nil {
		httpAgent = xhttp.New().SetHeader(xgin.CtxKey_TraceID, xstring.GetUUID())
	} else {
		httpAgent = ctx.InitRequest()
		hlog = ctx.Tech4GoodLog
	}
	return baseHttp(svr, url, method, expQuery, timeOut, headers, body, httpAgent, fill, errDetail, hlog)
}

// RasseHttp Rasse框架带上下文的请求
func RasseHttp(svr string, url string, method string, expQuery string, timeOut int, headers map[string]string,
	body string, ctx xcontext.XContext, fill bool, errDetail bool) (*xhttp.Res, error, *models.Response) {
	var hlog xlog.ILogger
	var httpAgent *xhttp.HTTPAgent
	if ctx == nil {
		httpAgent = xhttp.New().SetHeader(xgin.CtxKey_TraceID, xstring.GetUUID())
	} else {
		xhttp.NewWithContext(ctx)
	}
	return baseHttp(svr, url, method, expQuery, timeOut, headers, body, httpAgent, fill, errDetail, hlog)
}

// baseHttp 基础请求
func baseHttp(svr string, url string, method string, expQuery string,
	timeOut int, headers map[string]string, body string, httpAgent *xhttp.HTTPAgent,
	fill bool, errDetail bool, hlog xlog.ILogger) (*xhttp.Res, error, *models.Response) {
	var response models.Response
	fullUrl, token := getServiceUrl(svr)
	fullUrl = fullUrl + url
	if timeOut == 0 {
		timeOut = 10000 * 60
	}
	httpAgent.Url(fullUrl).Timeout(timeOut)
	if method != "" {
		httpAgent.Method(method)
	}
	if token != "" {
		httpAgent.SetRioToken(token)
	}
	if expQuery != "" {
		httpAgent.SetQuery("_exp", expQuery)
	}
	httpAgent.AddHeader("Content-Type", "application/json")
	if headers != nil {
		for key, value := range headers {
			httpAgent.AddHeader(key, value)
		}
	}
	jsonStr := []byte(body)
	httpAgent.SetBodyBytes(jsonStr)
	res, err := httpAgent.Request()
	if fill {
		fillResponse(fullUrl, res, err, hlog, &response, errDetail)
	}
	return res, err, &response
}

// getServiceUrl 根据服务名获取服务地址
func getServiceUrl(svcName string) (string, string) {
	if svcName != "" {
		svcs := viper.GetStringMapString("app.svcs")
		if url, ok := svcs[svcName]; ok {
			urls := strings.Split(url, ";")
			return urls[0], urls[1]
		}
	}
	return "", ""
}

// fillResponse 填充返回结构
func fillResponse(fullUrl string, res *xhttp.Res, err error, hlog xlog.ILogger,
	response *models.Response, errDetail bool) {
	if err != nil {
		hlog.Errorf(models.ErrorCode["SerError"].Log, fullUrl, err.Error())
		response.Code = models.ErrorCode["SerError"].Code
		response.ErrMsg = models.ErrorCode["SerError"].Msg
		if errDetail {
			response.ErrMsg = response.ErrMsg.(string) + err.Error()
		}
	} else {
		if res != nil && res.Body != nil {
			var json = jsoniter.ConfigCompatibleWithStandardLibrary
			err = json.Unmarshal(res.Body.Bytes(), &response)
			if err != nil {
				hlog.Errorf(models.ErrorCode["SerResJsonError"].Log, fullUrl, err.Error())
				response.Code = models.ErrorCode["SerResJsonError"].Code
				response.ErrMsg = models.ErrorCode["SerResJsonError"].Msg
				if errDetail {
					response.ErrMsg = response.ErrMsg.(string) + err.Error()
				}
			} else {
				if response.Code != 0 {
					hlog.Errorf(models.ErrorCode["SerError"].Log, fullUrl, response.ErrMsg.(string))
					response.Code = models.ErrorCode["SerError"].Code
					if errDetail {
						response.ErrMsg = models.ErrorCode["SerError"].Msg + response.ErrMsg.(string)
					} else {
						response.ErrMsg = models.ErrorCode["SerError"].Msg
					}
				}
			}
		} else {
			hlog.Errorf(models.ErrorCode["SerResEmpty"].Log, fullUrl)
			response.Code = models.ErrorCode["SerResEmpty"].Code
			response.ErrMsg = models.ErrorCode["SerResEmpty"].Msg
		}
	}
}
