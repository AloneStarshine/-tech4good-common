package mo

import (
	"time"
)

// Roles 功能表
type Roles struct {
	// CreateID column comments
	CreateID int `gorm:"column:create_id" json:"create_id"`
	// CreateName column comments
	CreateName string `gorm:"column:create_name" json:"create_name"`
	// CreateTime column comments
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// Desc 描述
	Desc string `gorm:"column:desc" json:"desc"`
	// Icon pc端图标
	Icon string `gorm:"column:icon" json:"icon"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// LastUpdateID column comments
	LastUpdateID int `gorm:"column:last_update_id" json:"last_update_id"`
	// LastUpdateName column comments
	LastUpdateName string `gorm:"column:last_update_name" json:"last_update_name"`
	// LastUpdateTime column comments
	LastUpdateTime time.Time `gorm:"column:last_update_time" json:"last_update_time"`
	// MobileIcon 移动端图标
	MobileIcon string `gorm:"primary_key;column:mobile_icon" json:"mobile_icon"`
	// Name 名称
	Name string `gorm:"column:name" json:"name"`
	// Status 可用状态
	Status int `gorm:"column:status" json:"status"`
	// Year 年份
	Year string `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (r *Roles) TableName() string {
	return "roles"
}
