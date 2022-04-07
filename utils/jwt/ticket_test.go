package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

//func TestName(t *testing.T) {
//	param := make(map[string]string)
//	param["sub_id"] = "1"
//	param["application_flag"] = "volunteer"
//	token, err := GetToken(param)
//	if err != nil {
//		fmt.Print(err)
//		return
//	}
//	fmt.Println(token)
//
//	parse, err := ParseToken(token)
//	fmt.Println(parse.SubId + "   " + parse.AppFlag)
//}

func TestGetUserToken(t *testing.T) {
	token := UserToken{
		StandardClaims: jwt.StandardClaims{},
		SubId:          "",
	}
	str, err := token.GetToken()
	if err != nil {
		return
	}
	fmt.Println(str)
}

func TestGetAppToken(t *testing.T) {
	Convey("TestGetAppToken", t, func() {
		token := AppToken{
			StandardClaims: jwt.StandardClaims{},
			AppFlag:        "volunteer",
			AppId:          "",
		}
		_, err := token.GetToken()
		So(err, ShouldBeNil)

	})
}

func TestGetToken(t *testing.T) {
	Convey("TestGetToken", t, func() {
		key := "testkey"
		Convey("case apptoken", func() {
			token := AppToken{
				StandardClaims: jwt.StandardClaims{},
				AppFlag:        "volunteer",
				AppId:          "",
			}
			str, err := GetToken(token, key)
			So(err, ShouldBeNil)
			So(str, ShouldNotEqual, "")
		})
		Convey("case usertoken", func() {
			token := UserToken{
				StandardClaims: jwt.StandardClaims{},
				SubId:          "",
			}
			str, err := GetToken(token, key)
			So(err, ShouldBeNil)
			So(str, ShouldNotEqual, "")
		})

	})
}

func TestParseUserToken(t *testing.T) {
	Convey("TestParseUserToken", t, func() {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWJfaWQiOiIifQ.1ALdijZlO9omey0fedlVrIuj6mqkS5YoIl8tbQB0T-s"
		res, err := ParseUserToken(token)
		So(err, ShouldBeNil)
		So(res.SubId, ShouldEqual, "")

	})
}

func TestParseAppToken(t *testing.T) {
	Convey("TestParseAppToken", t, func() {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhcHBsaWNhdGlvbl9mbGFnIjoidm9sdW50ZWVyIiwiYXBwX2lkIjoiIn0.8BhFwq_lnC4xWz8A_eY_GtSQDeczAk9GFimERgyE5oo"
		res, err := ParseAppToken(token)
		So(err, ShouldEqual, nil)
		So(res.AppId, ShouldEqual, "")

	})
}
