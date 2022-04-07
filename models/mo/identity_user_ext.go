package mo

import (
	"time"
)

// UserExt 用户扩展表
type UserExt struct {
	// CreateTime 添加时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 数据创建者
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// ObjName 用户其它额外信息健名
	ObjName string `gorm:"column:obj_name" json:"obj_name"`
	// ObjValue 用户其它额外信息健值
	ObjValue string `gorm:"column:obj_value" json:"obj_value"`
	// SubID 用于关联其它表的sub_id
	SubID string `gorm:"column:sub_id" json:"sub_id"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 数据更新者
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (u *UserExt) TableName() string {
	return "user_ext"
}
