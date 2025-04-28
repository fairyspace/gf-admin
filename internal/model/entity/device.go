// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure for table device.
type Device struct {
	Id         int         `json:"id"         orm:"id"          description:""`        //
	DeviceId   string      `json:"deviceId"   orm:"device_id"   description:"自定义设备ID"` // 自定义设备ID
	Mac        string      `json:"mac"        orm:"mac"         description:"mac地址"`   // mac地址
	Ip         string      `json:"ip"         orm:"ip"          description:"Ip地址"`    // Ip地址
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`    // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"`    // 更新时间
	ActiveTime *gtime.Time `json:"activeTime" orm:"active_time" description:"最后活跃时间"`  // 最后活跃时间
	Able       int         `json:"able"       orm:"able"        description:"是否启用"`    // 是否启用
	DeviceName string      `json:"deviceName" orm:"device_name" description:"设备名字"`    // 设备名字
	DeviceImg  string      `json:"deviceImg"  orm:"device_img"  description:"设备图片"`    // 设备图片
	Remark     string      `json:"remark"     orm:"remark"      description:""`        //
	UserId     int         `json:"userId"     orm:"user_id"     description:"绑定的用户"`   // 绑定的用户
}
