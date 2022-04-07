package utils

import (
	"fmt"
	"github.com/speps/go-hashids"

	//"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"unsafe"
	//"regexp"
	"testing"
)

var salt = "o23ejsd5"
var length = 10

func TestHashEncodeInt(t *testing.T) {
	Convey("TestHashEncodeInt", t, func() {
		init := &hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: length,
			Salt:      salt,
		}
		hashid, _ := hashids.NewWithData(init)
		myHash := &MyHash{hashid}
		Convey("case1", func() {
			num := 1035
			str := myHash.HashEncodeInt(num)
			So(str, ShouldEqual, "8YNpqDDbVJ") // 七彩石的实例化失败
		})
		Convey("case2", func() {
			num := 1033
			str := myHash.HashEncodeInt(num)
			So(str, ShouldEqual, "JYlg8d9gej")
		})

	})
}

func TestHashDecodeInt(t *testing.T) {
	Convey("TestHashDecodeInt", t, func() {
		init := &hashids.HashIDData{
			Alphabet:  hashids.DefaultAlphabet,
			MinLength: length,
			Salt:      salt,
		}
		hashid, _ := hashids.NewWithData(init)
		myHash := &MyHash{hashid}
		Convey("case1", func() {
			num, _ := myHash.HashDecodeInt("VmLpoE7w02") // 对加密过的字符串常见尝试解密
			num1, _ := myHash.HashDecodeInt("ERZbBjEbJ2") // 对加密过的字符串常见尝试解密
			fmt.Println(num,num1)
			So(num, ShouldEqual, 9)
		})
		Convey("case2", func() {
			num, _ := myHash.HashDecodeInt("JYlg8d9gej") // 对加密过的字符串常见尝试解密
			So(num, ShouldEqual, 1033)
		})
	})
}

type Book struct {
	Id   int    `json:"id" hashids:"true"`
	Name string `json:"name"`
}

type Person struct {
	Name string `json:"name" hashids:"true"`
	Age  int    `json:"age"`
}

func TestNewConfigWithHashIDs(t *testing.T) {
	Convey("NewConfigWithHashIDs", t, func() {
		Convey("case1", func() {
			result := NewConfigWithHashIDs("93hs", 10)
			bytes, _ := result.Marshal(Book{
				Id:   1234,
				Name: "Jane Eyre",
			})
			So(string(bytes), ShouldEqual, "{\"id\":\"XPZQx0JQnV\",\"name\":\"Jane Eyre\"}")

		})
		Convey("case2", func() {
			//验证hashid不能加密非int数据
			result := NewConfigWithHashIDs(salt, length)
			bytes, _ := result.Marshal(Person{
				Name: "xiao wang",
				Age:  30,
			})
			//fmt.Println(string(bytes))
			So(string(bytes), ShouldEqual, "{\"name\":\"xiao wang\",\"age\":30}")
		})
	})
}

func TestMyHash_HashDecodeInt(t *testing.T) {
	var typeName string
	var ptr unsafe.Pointer
	var val int64
	//var ErrNotInteger = errors.New("not integer")
	switch typeName {
	case "int":
		ip := (*int)(ptr)
		*ip = int(val)
	case "uint":
		ip := (*uint)(ptr)
		*ip = uint(val)
	case "int8":
		ip := (*int8)(ptr)
		*ip = int8(val)
	case "uint8":
		ip := (*uint8)(ptr)
		*ip = uint8(val)
	case "int16":
		ip := (*int16)(ptr)
		*ip = int16(val)
	case "uint16":
		ip := (*uint16)(ptr)
		*ip = uint16(val)
	case "int32":
		ip := (*int32)(ptr)
		*ip = int32(val)
	case "uint32":
		ip := (*uint32)(ptr)
		*ip = uint32(val)
	case "int64":
		ip := (*int64)(ptr)
		*ip = int64(val)
	case "uint64":
		ip := (*uint64)(ptr)
		*ip = uint64(val)
	default:
		return
	}
	err := setIntValue(typeName, ptr, val)
	fmt.Println(err)
	So(err, ShouldEqual, val)
	So(setIntValue, ShouldEqual, val)
}

//判断指针是否为空
//func TestFuncEncoder_IsEmpty(t *testing.T) {
//	var ptr unsafe.Pointer
//	encoder := &funcEncoder{}
//	bools := encoder.isEmptyFunc
//	So(bools,ShouldEqual,false)
//	bool := encoder.IsEmpty(ptr)
//	So(bool,ShouldEqual,true)
//
//}

//func TestFuncDecoder_Decode(t *testing.T) {
//	Convey("Decode",t, func() {
//		var ptr unsafe.Pointer
//		var iter *jsoniter.Iterator
//		var decoder *funcDecoder
//		gomonkey.ApplyFunc(decoder.Decode, func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {})
//		decoder.Decode(ptr,iter)
//		So(decoder.Decode,ShouldEqual,nil)
//		//decoder.Decode(ptr,iter)
//	})
//}
