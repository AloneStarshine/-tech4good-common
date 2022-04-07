package mo

import (
	"time"
)

// UserCompany 用户公司表
type UserCompany struct {
	// CompanyFlag 公司标识
	CompanyFlag string `gorm:"column:company_flag" json:"company_flag"`
	// CompanyName 公司名称
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 数据创建者
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// PersonID person表中的person_id
	PersonID string `gorm:"column:person_id" json:"person_id"`
	// SubID sub_id
	SubID string `gorm:"column:sub_id" json:"sub_id"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 数据更新者
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (u *UserCompany) TableName() string {
	return "user_company"
}
