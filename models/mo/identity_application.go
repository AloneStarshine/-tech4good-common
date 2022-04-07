package mo

import (
	"time"
)

// Application 应用信息表
type Application struct {
	// Appid appid
	Appid string `gorm:"column:appid" json:"appid"`
	// ApplicationFlag 应用标识，唯一
	ApplicationFlag string `gorm:"column:application_flag" json:"application_flag"`
	// ApplicationName 应用名称
	ApplicationName string `gorm:"column:application_name" json:"application_name"`
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 数据创建者
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// IsCompanyVerified 0不启用企业认证  1启用企业认证，企业认证指和对方企业对接信息
	IsCompanyVerified int `gorm:"column:is_company_verified" json:"is_company_verified"`
	// IsVerified 0不启用实名认证 1启用实名认证
	IsVerified int `gorm:"column:is_verified" json:"is_verified"`
	// LoginConfig 登录配置，使用数组键值对来匹配登录方式
	LoginConfig string `gorm:"column:login_config" json:"login_config"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// Token 应用令牌
	Token string `gorm:"column:token" json:"token"`
	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 数据修改者
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (a *Application) TableName() string {
	return "application"
}
