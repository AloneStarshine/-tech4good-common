package mo

import (
	"time"
)

// Person 用户实名表
type Person struct {
	// CreateTime 添加时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// CreateUser 数据添加人
	CreateUser string `gorm:"column:create_user" json:"create_user"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Idcard 身份证号，可以考虑加密
	Idcard string `gorm:"column:idcard" json:"idcard"`
	// PersonID 生成的user_id，唯一
	PersonID string `gorm:"column:person_id" json:"person_id"`
	// Sex 0女  1男  2保密
	Sex int `gorm:"column:sex" json:"sex"`
	// Summary 备注
	Summary string `gorm:"column:summary" json:"summary"`
	// Truename 真实姓名
	Truename string `gorm:"column:truename" json:"truename"`
	// UpdateTime 更新时间
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	// UpdateUser 数据修改人
	UpdateUser string `gorm:"column:update_user" json:"update_user"`
}

// TableName sets the insert table name for this struct type
func (p *Person) TableName() string {
	return "person"
}
