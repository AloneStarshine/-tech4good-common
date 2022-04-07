package jwt

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestRsa(t *testing.T) {
	Convey("TestRsa", t, func() {
		msg := "hello world"
		cypher, err := RsaEncrypt([]byte(msg))
		So(err, ShouldBeNil)
		plaint, err := RsaDecrypt(cypher)
		So(string(plaint), ShouldEqual, msg)
		So(err, ShouldBeNil)
	})

}
