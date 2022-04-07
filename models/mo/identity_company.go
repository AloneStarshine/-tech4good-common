package mo

import (
	"time"
)

// Company 公司表
type Company struct {
	// CompanyFlag 公司标识，唯一值
	CompanyFlag string `gorm:"column:company_flag" json:"company_flag"`
	// CompanyName 公司名称
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	// CreateTime 添加时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 数据添加者
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// UpdateTime 修改时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 数据修改者
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (c *Company) TableName() string {
	return "company"
}
