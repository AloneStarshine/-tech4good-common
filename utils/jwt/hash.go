package jwt

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
)

// MD5 md5大写hash
func MD5(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

// SHA256 获取字符串sha256结果
func SHA256(str string) string {
	h := sha256.New()
	h.Write([]byte(str))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

// MD516 获取16位md5结果
func MD516(str string) string {
	return MD5(str)[8:24]
}
