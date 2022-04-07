package jwt

import (
	"errors"
	"git.code.oa.com/fip-team/rasse/xlog"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
)

// UserToken 用户令牌ticket 数据
type UserToken struct {
	jwt.StandardClaims
	SubId           string `json:"sub_id"`
	ApplicationFlag string `json:"application_flag"`
	Timestamp       int64  `json:"timestamp"`
}

// AppToken 应用令牌token
type AppToken struct {
	jwt.StandardClaims
	AppFlag string `json:"application_flag"`
	AppId   string `json:"app_id"`
}

// GetToken 加密生成应用令牌
func (a *AppToken) GetToken() (token string, err error) {
	var key = viper.GetString("jwt.secret.application")
	token, err = GetToken(a, key)
	return
}

// GetToken 生成jwt令牌
func (u *UserToken) GetToken() (token string, err error) {
	var key = viper.GetString("jwt.secret.user")
	token, err = GetToken(u, key)
	return
}

// GetToken 生成令牌公共方法
func GetToken(c jwt.Claims, key string) (string, error) {
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	var secret = []byte(key)
	sign, err := token.SignedString(secret)
	if err != nil {
		xlog.Error(err)
		return "", err
	}
	return sign, nil
}

// ParseUserToken 解析用户令牌
func ParseUserToken(tokenString string) (*UserToken, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &UserToken{}, func(token *jwt.Token) (i interface{}, err error) {
		MySecret := []byte(viper.GetString("jwt.secret.user"))
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*UserToken); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// ParseAppToken 解析企业令牌
func ParseAppToken(tokenString string) (*AppToken, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &AppToken{}, func(token *jwt.Token) (i interface{}, err error) {
		MySecret := []byte(viper.GetString("jwt.secret.application"))
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*AppToken); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
