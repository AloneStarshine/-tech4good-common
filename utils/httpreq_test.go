package utils

import (
	"git.code.oa.com/fip-team/rasse/xcontext"
	"git.code.oa.com/tech4good/common/context"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// Http 带上下文的请求
func TestHttp(t *testing.T) {
	Convey("Http", t, func() {
		var svr, url, method, expQuery, body string
		var timeOut int
		var headers map[string]string
		var fill, errDetail bool
		var ctx *context.BaseContext
		_, err, _ := Http(svr, url, method, expQuery, timeOut, headers, body, ctx, fill, errDetail)
		So(err, ShouldNotBeNil)
	})
}

// RasseHttp Rasse框架带上下文的请求
func TestRasseHttp(t *testing.T) {
	Convey("RasseHttp", t, func() {
		var svr, url, method, expQuery, body string
		var timeOut int
		var headers map[string]string
		var fill, errDetail bool
		var ctx xcontext.XContext
		//aa:= xhttp.NewWithContext(ctx)
		//convey.So(aa,convey.ShouldBeIn,httpAgent)
		_, err, _ := RasseHttp(svr, url, method, expQuery, timeOut, headers, body, ctx, fill, errDetail)
		So(err, ShouldNotBeNil)
	})
}
