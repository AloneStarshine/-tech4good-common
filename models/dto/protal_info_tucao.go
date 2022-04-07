package dto

// TcxHookData 兔小巢hook数据
type TcxHookData struct {
	ID      string      `json:"id"`
	Type    string      `json:"type"`
	Payload PayloadData `json:"payload"`
}

// PayloadData Payload数据
type PayloadData struct {
	Post  interface{} `json:"post"`
	Reply interface{} `json:"reply"`
}

// PostData 推送数据
type PostData struct {
	ID            string      `json:"id"`
	HasAdminReply bool        `json:"has_admin_reply"`
	AvatarUrl     string      `json:"avatar_url"`
	Nickname      string      `json:"nick_name"`
	Content       string      `json:"content"`
	User          UserData    `json:"user"`
	IsAdmin       bool        `json:"is_admin"`
	RepliesAll    interface{} `json:"replies_all"`
}

// UserData 用户数据
type UserData struct {
	ID       int    `json:"id"`
	OpenId   string `json:"openid"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	IsAdmin  string `json:"is_admin"`
}

// ReplyData 反馈数据
type ReplyData struct {
	Id            string `json:"id"`
	FTitleId      string `json:"f_title_id"`
	Type          int    `json:"type"`
	ParentReplyId string `json:"parent_reply_id"`
	UserId        int    `json:"user_id"`
	NickName      string `json:"nick_name"`
	AvatarUrl     string `json:"avatar_url"`
	Content       string `json:"content"`
}
