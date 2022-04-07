package mo

import (
	"github.com/guregu/null"
	"time"
)

// Category 案例分类表
type Category struct {
	// CreateID column comments
	CreateID int `gorm:"column:create_id" json:"create_id"`
	// CreateName column comments
	CreateName string `gorm:"column:create_name" json:"create_name"`
	// CreateTime column comments
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Label 分类名称
	Label string `gorm:"column:label" json:"label"`
	// LastUpdateID column comments
	LastUpdateID int `gorm:"column:last_update_id" json:"last_update_id"`
	// LastUpdateName column comments
	LastUpdateName string `gorm:"column:last_update_name" json:"last_update_name"`
	// LastUpdateTime column comments
	LastUpdateTime time.Time `gorm:"column:last_update_time" json:"last_update_time"`
	// Status column comments
	Status int `gorm:"column:status" json:"status"`
	// ViewSort 展示顺序
	ViewSort null.Int `gorm:"column:view_sort" json:"view_sort"`
	// Year 年份
	Year string `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (c *Category) TableName() string {
	return "category"
}
