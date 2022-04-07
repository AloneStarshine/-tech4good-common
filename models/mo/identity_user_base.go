package mo

import (
	"time"
)

// UserBase 用户基础表
type UserBase struct {
	// ApplicationFlag 应用标识
	ApplicationFlag string `gorm:"column:application_flag" json:"application_flag"`
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
func (u *UserBase) TableName() string {
	return "user_base"
}
