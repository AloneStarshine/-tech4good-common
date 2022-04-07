package mo

import (
	"github.com/guregu/null"
	"time"
)

// AedDeviceInfo table comment
type AedDeviceInfo struct {
	// BatchNumber 数据批次号
	BatchNumber string `gorm:"column:batch_number" json:"batch_number"`
	// CityNumber aed设备所在城市编码；需关联城市表
	CityNumber string `gorm:"column:city_number" json:"city_number"`
	// ConnectedSystem aed设备是否连接其他急救系统;Y-已连接，N-未连接
	ConnectedSystem string `gorm:"column:connected_system" json:"connected_system"`
	// County aed设备所在市所属区/县
	County string `gorm:"column:county" json:"county"`
	// CreateTime 该条设备数据创建时间
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// DataProvider 设备数据提供者
	DataProvider string `gorm:"column:data_provider" json:"data_provider"`
	// DataSourceType 设备数据来源方式
	DataSourceType string `gorm:"column:data_source_type" json:"data_source_type"`
	// DeviceAddress aed设备地址
	DeviceAddress string `gorm:"column:device_address" json:"device_address"`
	// DeviceAvailableTime aed设备所在地可取用时段
	DeviceAvailableTime string `gorm:"column:device_available_time" json:"device_available_time"`
	// DeviceBrand aed设备生产商
	DeviceBrand string `gorm:"column:device_brand" json:"device_brand"`
	// DeviceLatitude 纬度
	DeviceLatitude string `gorm:"column:device_latitude" json:"device_latitude"`
	// DeviceLongitude 经度
	DeviceLongitude string `gorm:"column:device_longitude" json:"device_longitude"`
	// DeviceMaintenancer aed设备维护方
	DeviceMaintenancer string `gorm:"column:device_maintenancer" json:"device_maintenancer"`
	// DeviceNumber aed设备管理编号，SZ0755-序列号
	DeviceNumber string `gorm:"column:device_number" json:"device_number"`
	// DeviceOwner aed设备所属方
	DeviceOwner string `gorm:"column:device_owner" json:"device_owner"`
	// DeviceSerialNumber aed设备自身的序列号
	DeviceSerialNumber string `gorm:"column:device_serial_number" json:"device_serial_number"`
	// DeviceStatus aed设备状态;Y-正常可用，N-不可用
	DeviceStatus string `gorm:"column:device_status" json:"device_status"`
	// EmergencyName aed设备紧急联系人
	EmergencyName null.String `gorm:"column:emergency_name" json:"emergency_name"`
	// EmergencySystemName aed设备连接的其他急救系统名称
	EmergencySystemName null.String `gorm:"column:emergency_system_name" json:"emergency_system_name"`
	// EmergencyTel aed设备紧急联系电话
	EmergencyTel null.String `gorm:"column:emergency_tel" json:"emergency_tel"`
	// ID column comments
	ID int `gorm:"primary_key;column:id" json:"id"`
	// LastUpdateTime 上次更新数据时间
	LastUpdateTime time.Time `gorm:"column:last_update_time" json:"last_update_time"`
	// MaintenancePeriod 设备维护周期
	MaintenancePeriod null.Int `gorm:"column:maintenance_period" json:"maintenance_period"`
	// ServiceActualTime 设备使用实际终止日期
	ServiceActualTime time.Time `gorm:"column:service_actual_time" json:"service_actual_time"`
	// ServiceExpireTime 设备使用寿命理论终止日期
	ServiceExpireTime time.Time `gorm:"column:service_expire_time" json:"service_expire_time"`
	// ServiceStartTime 设备投入使用开始日期
	ServiceStartTime time.Time `gorm:"column:service_start_time" json:"service_start_time"`
	// Street aed设备所在街道
	Street string `gorm:"column:street" json:"street"`
}

// TableName sets the insert table name for this struct type
func (a *AedDeviceInfo) TableName() string {
	return "aed_device_info"
}
