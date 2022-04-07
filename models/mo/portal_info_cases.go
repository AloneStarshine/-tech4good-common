package mo

import (
	"github.com/guregu/null"
	"time"
)

// Cases 案例表
type Cases struct {
	// AbilityID 能力id
	AbilityID int `gorm:"column:ability_id" json:"ability_id"`
	// Application 技术运用描述
	Application string `gorm:"column:application" json:"application"`
	// CategoryID 案例分类id
	CategoryID int `gorm:"column:category_id" json:"category_id"`
	// Contacts 联系方式
	Contacts string `gorm:"column:contacts" json:"contacts"`
	// CreateID column comments
	CreateID int `gorm:"column:create_id" json:"create_id"`
	// CreateName column comments
	CreateName string `gorm:"column:create_name" json:"create_name"`
	// CreateTime column comments
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// Desc 案例描述
	Desc string `gorm:"column:desc" json:"desc"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// Introduce 简介
	Introduce string `gorm:"column:introduce" json:"introduce"`
	// IsHomeView 是否在首页展示：1-展示，0-不展示
	IsHomeView int `gorm:"column:is_home_view" json:"is_home_view"`
	// LastUpdateID column comments
	LastUpdateID int `gorm:"column:last_update_id" json:"last_update_id"`
	// LastUpdateName column comments
	LastUpdateName string `gorm:"column:last_update_name" json:"last_update_name"`
	// LastUpdateTime column comments
	LastUpdateTime time.Time `gorm:"column:last_update_time" json:"last_update_time"`
	// MobileDetailImgURL 移动端案列简介图片地址
	MobileDetailImgURL null.String `gorm:"column:mobile_detail_img_url" json:"mobile_detail_img_url"`
	// MobileHomeImgURL 移动端案列首页图片地址
	MobileHomeImgURL null.String `gorm:"column:mobile_home_img_url" json:"mobile_home_img_url"`
	// MobileListImgURL 移动端案列列表图片地址
	MobileListImgURL null.String `gorm:"column:mobile_list_img_url" json:"mobile_list_img_url"`
	// PcDetailImgURL 案例简介图片地址
	PcDetailImgURL null.String `gorm:"column:pc_detail_img_url" json:"pc_detail_img_url"`
	// PcHomeImgURL 案例首页图片地址
	PcHomeImgURL null.String `gorm:"column:pc_home_img_url" json:"pc_home_img_url"`
	// PcListImgURL 案列列表图片地址
	PcListImgURL null.String `gorm:"column:pc_list_img_url" json:"pc_list_img_url"`
	// Significance 项目意义
	Significance string `gorm:"column:significance" json:"significance"`
	// Status 案例状态0-下架，1-正常
	Status int `gorm:"column:status" json:"status"`
	// Title 案例名称
	Title string `gorm:"column:title" json:"title"`
	// ViewSort 首页展示顺序
	ViewSort int `gorm:"column:view_sort" json:"view_sort"`
	// Website 网站
	Website string `gorm:"column:website" json:"website"`
	// Year 年份
	Year string `gorm:"column:year" json:"year"`
}

// TableName sets the insert table name for this struct type
func (c *Cases) TableName() string {
	return "cases"
}
