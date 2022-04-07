package mo

import (
	"time"
)

// Department 部门表
type Department struct {
	// CompanyFlag 公司标识
	CompanyFlag string `gorm:"column:company_flag" json:"company_flag"`
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 创建数据者
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// DepartmentName 部门名称
	DepartmentName string `gorm:"column:department_name" json:"department_name"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 更新数据者
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (d *Department) TableName() string {
	return "department"
}
