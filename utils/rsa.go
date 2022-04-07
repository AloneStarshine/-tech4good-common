package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

//密钥为pem格式

//公钥加密，输入string格式的明文，输出hex string格式的密文
func RsaEncrypt(msg string, key string) (string, error) {

	//解析公钥
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)

	//加密
	cyphertext, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(msg))
	if err != nil {
		return "", err
	}

	//返回hex string格式的密文
	return hex.EncodeToString(cyphertext), nil
}

//私钥解密，输入hex string格式的密文，输出string格式的明文
func RsaDecrypt(cypher string, key string) (string, error) {
	if cypher == "" {
		return "", nil
	}

	//hex string转[]byte
	msg, err := hex.DecodeString(cypher)
	if err != nil {
		return "", err
	}

	//解析私钥
	block, _ := pem.Decode([]byte(key))
	if block == nil {
		return "", errors.New("empty key")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	//解密
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, msg)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
