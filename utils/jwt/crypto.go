package jwt

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// 私钥生成
//openssl genrsa -out rsa_private_key.pem 1024
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQD6eb7uCpIQXIg4R7na5bh6u9nSO4Zp4S7AbKpgv+EvngTGng3P
/rhTgxfv0yrr7w6YZSv6JtaXkCCNfYP/BP7PhEevTXH5DBjp0wc9GCTTc3MnPON3
NIlQb7gXXQLx4FkSdNA+hgdrG3s/9oiTVT0iAIovKHKYU1ol8ZceXrwcUwIDAQAB
AoGAQndyKhrV/c+AOmcWM7dICBG3UKmJFqmxzVBIuhnQ+ODW5Znlkm9GnKqp/HMt
7aPnXJtkWyJZSajuan2HPHIn54UO5p/XN4WGwiBcw0rxi3xYdu6DdUrAZKqOVeTQ
MHOH885n4RYhJsMm7RyuaBHGgqItzXZZ97NKfMNQkGd1xIECQQD+Ca4FFbrhK9LY
MWWnG5Nd2K0eJ9PLhIMMN0aG3cuCkyfmuhBlc0wM0utun5j9BJ1wvnm9CCnvANEq
/6i0XAPBAkEA/GkFs9mYYJAIjVzFxE+fHrdLky3e0W8mZx/VA85eu/oQrNtKuoUB
VSSJLSMn8f6pQ42YLtIiec7uNTlqgdkVEwJBAKKN7xyx2vNa54APm8xiiNn0XFJ/
ibchA/o9JJQIOMFFCLNLPFKuhGtwS9Ztqae93EDYoW2kW7DkBPROw9UlTAECQQCS
whmPta/UTUq7rrpKZyUUfeySObR5P1Ar26VGHkKUt1PkvWhYxKa+s4yS0wMRwEj4
PybB6mojOr7j8WtM7kRDAkBO37AjTtdzVxX5p2Hnd76crr5vbRfXYH0YMcCVHPEY
ZwzqkiW0RzdvBYArm7ABXodxXYPG0x0Kit8i1sPJF1kB
-----END RSA PRIVATE KEY-----
`)

// 公钥: 根据私钥生成
//openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQD6eb7uCpIQXIg4R7na5bh6u9nS
O4Zp4S7AbKpgv+EvngTGng3P/rhTgxfv0yrr7w6YZSv6JtaXkCCNfYP/BP7PhEev
TXH5DBjp0wc9GCTTc3MnPON3NIlQb7gXXQLx4FkSdNA+hgdrG3s/9oiTVT0iAIov
KHKYU1ol8ZceXrwcUwIDAQAB
-----END PUBLIC KEY-----
`)

// RsaEncrypt 加密
func RsaEncrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
