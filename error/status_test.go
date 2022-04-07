package error

import (
	_ "errors"
	_ "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCode2(t *testing.T) {
	Convey("error", t, func() {
		Convey("GetCode2", func() {
			err := New(OtherError).Err()
			if err == nil {
				str := GetCode(err)
				So(str, ShouldBeNil, OtherError)
			}
		})
		Convey("Newf", func() {
			var code Code = 200006
			sa := Newf(code, "%s login fail!", "zhangsan")
			So(sa, ShouldNotBeIn, code)
		})

	})
}
