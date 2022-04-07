package mo

import (
	"time"
)

// UserDepartment 用户部门表
type UserDepartment struct {
	// CompanyFlag 公司标识
	CompanyFlag string `gorm:"column:company_flag" json:"company_flag"`
	// CompanyName 公司名称
	CompanyName string `gorm:"column:company_name" json:"company_name"`
	// CreateTime 添加时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 数据添加人
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// DepartmentName 部门名称
	DepartmentName string `gorm:"column:department_name" json:"department_name"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// PersonID 对应自然人的person_id
	PersonID string `gorm:"column:person_id" json:"person_id"`
	// SubID sub_id
	SubID string `gorm:"column:sub_id" json:"sub_id"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 数据更新人
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (u *UserDepartment) TableName() string {
	return "user_department"
}
