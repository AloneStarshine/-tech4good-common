package mo

import (
	"time"
)

// Home 首页表
type Home struct {
	// CreateID column comments
	CreateID int `gorm:"column:create_id" json:"create_id"`
	// CreateName column comments
	CreateName string `gorm:"column:create_name" json:"create_name"`
	// CreateTime column comments
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// Desc 网站简介
	Desc string `gorm:"column:desc" json:"desc"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// LastUpdateID column comments
	LastUpdateID int `gorm:"column:last_update_id" json:"last_update_id"`
	// LastUpdateName column comments
	LastUpdateName string `gorm:"column:last_update_name" json:"last_update_name"`
	// LastUpdateTime column comments
	LastUpdateTime time.Time `gorm:"column:last_update_time" json:"last_update_time"`
	// Status 可用状态1=可用,-1=不可用
	Status int `gorm:"column:status" json:"status"`
	// Year 年份
	Year string `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (h *Home) TableName() string {
	return "home"
}
