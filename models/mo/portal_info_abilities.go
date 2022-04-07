package mo

import (
	"github.com/guregu/null"
	"time"
)

// Abilities 能力表
type Abilities struct {
	// Contacts 联系方式
	Contacts string `gorm:"column:contacts" json:"contacts"`
	// CreateID column comments
	CreateID int `gorm:"column:create_id" json:"create_id"`
	// CreateName column comments
	CreateName string `gorm:"column:create_name" json:"create_name"`
	// CreateTime column comments
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// Desc 描述
	Desc string `gorm:"column:desc" json:"desc"`
	// Icon column comments
	Icon string `gorm:"column:icon" json:"icon"`
	// IconHover column comments
	IconHover string `gorm:"column:icon_hover" json:"icon_hover"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Introduce 简介
	Introduce string `gorm:"column:introduce" json:"introduce"`
	// IsHomeView 首页是否展示：0-不展示，1-展示
	IsHomeView int `gorm:"column:is_home_view" json:"is_home_view"`
	// LastUpdateID column comments
	LastUpdateID int `gorm:"column:last_update_id" json:"last_update_id"`
	// LastUpdateName column comments
	LastUpdateName string `gorm:"column:last_update_name" json:"last_update_name"`
	// LastUpdateTime column comments
	LastUpdateTime time.Time `gorm:"column:last_update_time" json:"last_update_time"`
	// Name 能力名称
	Name string `gorm:"column:name" json:"name"`
	// Status 是否有效；1-有效，0-无效
	Status null.Int `gorm:"column:status" json:"status"`
	// ViewSort 展示序号
	ViewSort int `gorm:"column:view_sort" json:"view_sort"`
	// Website 网站
	Website string `gorm:"column:website" json:"website"`
	// Year 年份
	Year string `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (a *Abilities) TableName() string {
	return "abilities"
}
