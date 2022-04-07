package dto

// AccessToken 凭证定义
type AccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// Param 微信参数
type Param struct {
	ToUser      string          `json:"touser"`
	TemplateId  string          `json:"template_id"`
	Url         string          `json:"url"`
	Miniprogram MiniprogramData `json:"miniprogram"`
	Data        PushData        `json:"data"`
}

// MiniprogramData 小程序数据
type MiniprogramData struct {
	Appid    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

// PushData 推送数据
type PushData struct {
	First    ValueData `json:"first"`
	Keyword1 ValueData `json:"keyword1"`
	Keyword2 ValueData `json:"keyword2"`
	Keyword3 ValueData `json:"keyword3"`
	Remark   ValueData `json:"remark"`
}

// ValueData 值数据
type ValueData struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

// Result 结束定义
type Result struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// OpenIDParam OpenID参数
type OpenIDParam struct {
	AppID  string `json:"appid"`
	Secret string `json:"secret"`
	Code   string `json:"code"`
}

// OpenIdResult OpenID结果
type OpenIdResult struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenId       string `json:"openid"`
	Scope        string `json:"scope"`
}

// UserInfo 用户信息
type UserInfo struct {
	OpenId     string `json:"openid"`
	Nickname   string `json:"nickname"`
	HeadImgUrl string `json:"headimgurl"`
}
