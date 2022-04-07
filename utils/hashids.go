package utils

import (
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
	"github.com/speps/go-hashids"
	"github.com/spf13/viper"
	"reflect"
	"unsafe"
)

// 该组件基于jsoniter
// 在mo或者dto中加入tag：hashid:"true"即可在序列化中将id进行混淆
// 序列化方法基于json-iterator，实现了encoder接口和decoder接口
// 加密算法基于go-hashids，是一种对称加密算法，只支持int类型
// 该算法可正逆向转化id
// 后续如果有别的加密可以基于此框架进行改造，放入json包中

var ErrNotInteger = errors.New("not integer")

// NewConfigWithHashIDs 初始化json序列化api
//  @param salt 密码加盐
//  @param minLength 密码长度
//  @return jsoniter.API 序列化API
//
//  [Example]
//  type Book struct {
// 	  	Id    int    `json:"id" hashids:"true"`
// 	  	Name  string `json:"name"`
//   }
//  var json = jsonhashids.NewConfigWithHashIDs("93hs", 10)
//  bytes, _ := json.Marshal(Book {
//  	Id:          1234,
//  	Name:        "Jane Eyre",
//  })
//  输出: {"id":"gYEL5rKBnd","name":"Jane Eyre"}
func NewConfigWithHashIDs(salt string, minLength int) jsoniter.API {

	e := NewHashIDsExtension(salt, minLength)
	config := jsoniter.ConfigCompatibleWithStandardLibrary
	config.RegisterExtension(e)

	return config
}

// MyHash 自定义 hash 加密
type MyHash struct {
	*hashids.HashID
}

// NewMyHash 根据七彩石配置创建 MyHash 实例
func NewMyHash() *MyHash {
	init := &hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		MinLength: viper.GetInt("hashid.length"),
		Salt:      viper.GetString("hashid.salt"),
	}
	hashid, _ := hashids.NewWithData(init)
	return &MyHash{hashid}
}

// HashEncodeInt 根据 hashID 规则 int covert to string
//  [Example]
//  myHash := utils.NewMyHash()
//  num := 1035
//  str := myHash.HashEncodeInt(num)
//  输出：str = 8YNpqDDbVJ
func (m *MyHash) HashEncodeInt(number int) string {
	res, err := m.Encode([]int{number})
	if err != nil {
		return ""
	}
	return res
}

// HashDecodeInt 根据 hashID 规则 string covert to int，返回第一个数
//  [Example]
//  myHash := utils.NewMyHash()
//  hash := 8YNpqDDbVJ
//  num, _ := myHash.HashDecodeInt(hash)
//  输出：num = 1035
func (m *MyHash) HashDecodeInt(hash string) (int, error) {
	decodeInt, err := m.DecodeWithError(hash)
	return decodeInt[0], err
}

// NewHashIDsExtension 对id进行加盐
func NewHashIDsExtension(salt string, minLength int) *HashIDsExtension {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = minLength
	h, _ := hashids.NewWithData(hd)
	return &HashIDsExtension{
		hashid: h,
	}
}

// HashIDsExtension hashId 扩展属性
type HashIDsExtension struct {
	hashid *hashids.HashID
	jsoniter.DummyExtension
}

//UpdateStructDescriptor 更新结构体属性
func (extension *HashIDsExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {

	for _, binding := range structDescriptor.Fields {

		tag := binding.Field.Tag().Get("hashids")
		if tag != "true" {
			continue
		}

		// 只支持int类型
		switch binding.Field.Type().Kind() {
		case reflect.Int:
		case reflect.Uint:
		case reflect.Int8:
		case reflect.Uint8:
		case reflect.Int16:
		case reflect.Uint16:
		case reflect.Int32:
		case reflect.Uint32:
		case reflect.Int64:
		case reflect.Uint64:
		default:
			continue
		}

		typeName := binding.Field.Type().String()
		binding.Encoder = &funcEncoder{fun: func(ptr unsafe.Pointer, stream *jsoniter.Stream) {

			i, err := int64Val(typeName, ptr)
			if err != nil {
				stream.Error = err
				return
			}

			// 后续可以替换算法
			hashed, err := extension.hashid.EncodeInt64([]int64{i})
			if err != nil {
				stream.Error = err
				return
			}
			stream.WriteString(hashed)
		}}

		binding.Decoder = &funcDecoder{fun: func(ptr unsafe.Pointer, iter *jsoniter.Iterator) {

			str := iter.ReadString()
			if str == "" {
				return
			}

			ints := extension.hashid.DecodeInt64(str)
			if len(ints) != 1 {
				return
			}

			setIntValue(typeName, ptr, ints[0])
		}}
	}
}

// funcDecoder 组合jsoniter.DecoderFunc方法
type funcDecoder struct {
	fun jsoniter.DecoderFunc
}

// Decode 反序列化
func (decoder *funcDecoder) Decode(ptr unsafe.Pointer, iter *jsoniter.Iterator) {
	decoder.fun(ptr, iter)
}

// funcDecoder 组合jsoniter.EncoderFunc方法
type funcEncoder struct {
	fun         jsoniter.EncoderFunc
	isEmptyFunc func(ptr unsafe.Pointer) bool
}

// Encode 实现组合jsoniter序列化方法
func (encoder *funcEncoder) Encode(ptr unsafe.Pointer, stream *jsoniter.Stream) {
	encoder.fun(ptr, stream)
}

// IsEmpty 判断指针是否为空
func (encoder *funcEncoder) IsEmpty(ptr unsafe.Pointer) bool {
	if encoder.isEmptyFunc == nil {
		return false
	}
	return encoder.isEmptyFunc(ptr)
}

// int64Val 转int64
func int64Val(typeName string, ptr unsafe.Pointer) (int64, error) {

	var i int64

	switch typeName {
	case "int":
		ip := (*int)(ptr)
		i = int64(*ip)
	case "uint":
		ip := (*uint)(ptr)
		i = int64(*ip)
	case "int8":
		ip := (*int8)(ptr)
		i = int64(*ip)
	case "uint8":
		ip := (*uint8)(ptr)
		i = int64(*ip)
	case "int16":
		ip := (*int16)(ptr)
		i = int64(*ip)
	case "uint16":
		ip := (*uint16)(ptr)
		i = int64(*ip)
	case "int32":
		ip := (*int32)(ptr)
		i = int64(*ip)
	case "uint32":
		ip := (*uint32)(ptr)
		i = int64(*ip)
	case "int64":
		ip := (*int64)(ptr)
		i = *ip
	case "uint64":
		ip := (*uint64)(ptr)
		i = int64(*ip)
	default:
		return 0, ErrNotInteger
	}

	return i, nil
}

// setIntValue 转int类型
func setIntValue(typeName string, ptr unsafe.Pointer, val int64) error {
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
		return ErrNotInteger
	}

	return nil
}
