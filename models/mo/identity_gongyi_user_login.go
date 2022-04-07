package mo

import (
	"time"
)

// GongyiUserLogin 公益小程序用户登录表
type GongyiUserLogin struct {
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 数据创建者
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// Email 邮箱
	Email string `gorm:"column:email" json:"email"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// LastLogintime 最后登录时间
	LastLogintime time.Time `gorm:"column:last_logintime" json:"last_logintime"`
	// LoginName 登录用户名
	LoginName string `gorm:"column:login_name" json:"login_name"`
	// LoginPassword 登录密码
	LoginPassword string `gorm:"column:login_password" json:"login_password"`
	// LoginType 登录方式
	LoginType string `gorm:"column:login_type" json:"login_type"`
	// Mobile 手机号
	Mobile string `gorm:"column:mobile" json:"mobile"`
	// QqUnionid qq登录的unionid
	QqUnionid string `gorm:"column:qq_unionid" json:"qq_unionid"`
	// Salt 用于加密的salt
	Salt string `gorm:"column:salt" json:"salt"`
	// SubID 返回给应用的sub_id
	SubID string `gorm:"column:sub_id" json:"sub_id"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 数据更新者
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
	// WxUnionid 微信登录的unionid
	WxUnionid string `gorm:"column:wx_unionid" json:"wx_unionid"`
}

// TableName sets the insert table name for this struct type
func (g *GongyiUserLogin) TableName() string {
	return "gongyi_user_login"
}
