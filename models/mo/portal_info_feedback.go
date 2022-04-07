package mo

import (
	"github.com/guregu/null"
)

// Feedback table comment
type Feedback struct {
	// ID id
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	// Mobile 手机
	Mobile string `gorm:"column:mobile" json:"mobile"`
	// FullName 姓名
	FullName null.String `gorm:"column:full_name" json:"full_name"`
	// Content 反馈内容
	Content string `gorm:"column:content" json:"content"`
}

// TableName sets the insert table name for this struct type
func (f *Feedback) TableName() string {
	return "feedback"
}
