package utils

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"regexp"
	"testing"
)

func isArrayEqual(a, b []string) bool {
	fmt.Println("actual", a)
	fmt.Println("expect", b)
	fmt.Println("-------------------")
	if len(a) != len(b) {
		return false
	}
	var lenmin int
	if len(a) < len(b) {
		lenmin = len(a)
	} else {
		lenmin = len(b)
	}
	for i := 0; i < lenmin; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestSplit(t *testing.T) {
	Convey("TestSplit", t, func() {
		Convey("case1", func() {
			s := "a,b,c"
			res := Split(s, "")
			So(isArrayEqual(res, []string{s}), ShouldBeTrue)
		})
		Convey("case2", func() {
			s := "a,b,c"
			res := Split(s)
			So(isArrayEqual(res, []string{s}), ShouldBeTrue)
		})
		Convey("case3", func() {
			s := "a,b,c"
			res := Split(s, ",")
			So(isArrayEqual(res, []string{"a", "b", "c"}), ShouldBeTrue)
		})
		Convey("case4", func() {
			s := "a,b;c"
			res := Split(s, ",", ";")
			So(isArrayEqual(res, []string{"a", "b", "c"}), ShouldBeTrue)
		})
		Convey("case5", func() {
			s := "a,b;c"
			res := Split(s, ";")
			So(isArrayEqual(res, []string{"a,b", "c"}), ShouldBeTrue)
		})
		Convey("case6", func() {
			s := "a,b,c"
			res := Split(s, "", ",")
			So(isArrayEqual(res, []string{s}), ShouldBeTrue)
		})
		Convey("case7", func() {
			s := "a,b;c"
			res := Split(s, ";", "", ",")
			So(isArrayEqual(res, []string{"a,b", "c"}), ShouldBeTrue)
		})
	})
}

func TestGetUUID(t *testing.T) {
	Convey("TestGetUUID", t, func() {
		Convey("case1", func() {
			result := GetUUID()
			match, _ := regexp.Match("-", []byte(result))
			So(match, ShouldBeFalse)
		})
	})
}
