package jwt

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestMD5(t *testing.T) {
	Convey("TestMD5", t, func() {
		str := MD5("hello")
		So(str, ShouldNotBeNil)
	})
}

func TestSHA256(t *testing.T) {
	Convey("TestSHA256", t, func() {
		str := "hello world\n"
		res := SHA256(str)
		So(res, ShouldEqual, "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447")
	})
}

func TestMD516(t *testing.T) {
	Convey("TestMD516", t, func() {
		str := MD516("hello")
		So(str, ShouldNotBeNil)
	})
}
