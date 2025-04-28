// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// License is the golang structure for table license.
type License struct {
	Id           int         `json:"id"           orm:"id"            description:""`        //
	LicenseId    string      `json:"licenseId"    orm:"license_id"    description:"许可证ID"`   // 许可证ID
	DeviceId     string      `json:"deviceId"     orm:"device_id"     description:"设备ID"`    // 设备ID
	DeviceName   string      `json:"deviceName"   orm:"device_name"   description:"设备名字"`    // 设备名字
	StartTime    *gtime.Time `json:"startTime"    orm:"start_time"    description:"开始时间"`    // 开始时间
	EndTime      *gtime.Time `json:"endTime"      orm:"end_time"      description:"结束时间"`    // 结束时间
	CreateTime   *gtime.Time `json:"createTime"   orm:"create_time"   description:"创建时间"`    // 创建时间
	UpdateTime   *gtime.Time `json:"updateTime"   orm:"update_time"   description:"更新时间"`    // 更新时间
	Type         string      `json:"type"         orm:"type"          description:"许可证订阅类型"` // 许可证订阅类型
	PayImg       string      `json:"payImg"       orm:"pay_img"       description:"支付截图"`    // 支付截图
	PayType      string      `json:"payType"      orm:"pay_type"      description:"支付类型"`    // 支付类型
	Money        float64     `json:"money"        orm:"money"         description:"金钱"`      // 金钱
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`      // 备注
	SerialNumber string      `json:"serialNumber" orm:"serial_number" description:"流水号"`     // 流水号
	Able         int         `json:"able"         orm:"able"          description:"是否启用状态"`  // 是否启用状态
}
