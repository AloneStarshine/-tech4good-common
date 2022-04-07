package dto

// 附近aed设备信息
type NearbyAedInfo struct {
	AedId              int    `json:"aed_id"`              //数据id
	SerialNumber       string `json:"serial_number"`       //设备序列号
	DeviceNumber       string `json:"device_number"`       //设备编号
	Address            string `json:"address"`             //地址
	Longitude          string `json:"longitude"`           //经度
	Latitude           string `json:"latitude"`            //纬度
	CityNumber         string `json:"city_number"`         //市代号
	CityName           string `json:"city_name"`           //市名称
	County             string `json:"county"`              //区县
	Street             string `json:"street"`              //街道、镇
	Brand              string `json:"brand"`               //生产商
	Owner              string `json:"owner"`               //设备所属方
	Maintenancer       string `json:"maintenancer"`        //维护方
	AvailableTime      string `json:"available_time"`      //aed设备所在地可取用时间段
	Status             string `json:"status"`              //aed设备状态
	IsConnected        string `json:"is_connected"`        //aed是否连接其他急救系统
	SystemName         string `json:"system_name"`         //aed连接其他急救系统名称
	EmergencyName      string `json:"emergency_name"`      //aed设备紧急联系人
	EmergencyTel       string `json:"emergency_tel"`       //aed设备紧急联系电话
	MaintenancerPeriod string `json:"maintenancer_period"` //设备维护周期
	ServiceStart       string `json:"service_start"`       //投入使用开始时间
	ServiceExpire      string `json:"service_expire"`      //使用寿命终止时间
	ServiceActual      string `json:"service_actual"`      //实际使用终止时间
}
