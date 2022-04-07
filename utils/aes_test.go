package utils

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAesCtrEncrypt(t *testing.T) {
	Convey("TestAesCtrEncrypt", t, func() {
		//测试加密空字符串
		Convey("encrypt empty string", func() {
			result, err := AesCtrEncrypt("")
			So(err, ShouldNotBeNil)
			So(result, ShouldEqual, "")
		})
		//测试解密空字符串
		Convey("decrypt empty string", func() {
			result, err := AesCtrDecrypt("")
			So(err, ShouldNotBeNil)
			So(result, ShouldEqual, "")
		})
		//测试加密普通字符串
		Convey("normal", func() {
			plaintext := "this is aes message"
			cypher, err := AesCtrEncrypt(plaintext)
			fmt.Println("cypher text", cypher)
			So(err, ShouldNotBeNil)
			So(cypher, ShouldEqual, "")
			decode, err := AesCtrDecrypt(cypher)
			So(err, ShouldNotBeNil)
			//		fmt.Println("decode text", decode)
			So(decode, ShouldEqual, "")
			So(plaintext, ShouldEqual, plaintext)
		})
		//测试加密unicode字符串
		Convey("Unicode string", func() {
			plaintext := "\uffff\u0fff\ufd00\ufffd\ufdfd\u1111\uf112\u1231\uffaa\ufffa"
			cypher, err := AesCtrEncrypt(plaintext)
			fmt.Println("cypher text", cypher)
			So(err, ShouldNotBeNil)
			decode, err := AesCtrDecrypt(cypher)
			So(err, ShouldNotBeNil)
			fmt.Println("decode text", decode)
			So(plaintext, ShouldEqual, plaintext)
		})
	})
}
