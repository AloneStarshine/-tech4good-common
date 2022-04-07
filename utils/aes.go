package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"github.com/spf13/viper"
)

//支持16、24、32bytes的密钥
//aes加密，输入string格式的明文，返回hex string格式的密文
func AesCtrEncrypt(plaintext string) (string, error) {
	key := []byte(viper.GetString("aes.key"))
	msg := []byte(plaintext)
	//	创建aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//	创建分组模式ctr
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	//	加密
	dst := make([]byte, len(msg))
	stream.XORKeyStream(dst, msg)

	//返回hex string格式的密文
	return hex.EncodeToString(dst), nil
}

//aes解密，输入hex string格式的密文，返回string格式的明文
func AesCtrDecrypt(ciphertext string) (string, error) {
	key := []byte(viper.GetString("aes.key"))
	msg, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	//	创建aes密码接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//	创建分组模式ctr
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)

	//	解密
	dst := make([]byte, len(msg))
	stream.XORKeyStream(dst, msg)

	//返回string格式的密文
	return string(dst), nil
}
