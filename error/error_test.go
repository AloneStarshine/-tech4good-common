package error

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

func TestError_Error(t *testing.T) {
	Convey("Error", t, func() {
		Convey("Error", func() {
			var code Code = 200000
			var status = &Status{Code: code, Message: "message"}
			err := Error{
				Status: status,
			}
			str := err.Error()
			So(str, ShouldNotBeNil)
		})
		Convey("GetCode", func() {
			_, errs := strconv.Atoi("error")
			s, ok := FromError(errs)
			if ok {
				fmt.Println(s.Code)
				fmt.Println(s.Message)
			}
			str := GetCode(errs)
			So(str, ShouldNotBeNil)
		})
		Convey("GetMessage", func() {
			var code Code = -1
			var status = &Status{Code: code, Message: "缺少参数"}
			var err error = &Error{
				Status: status,
			}
			s, ok := FromError(err)
			if ok {
				fmt.Println(s.Code)
				fmt.Println(s.Message)
			}
			str := GetMessage(err)
			So(str, ShouldNotBeNil)
		})

	})
}
