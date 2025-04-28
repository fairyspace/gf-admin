// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Device is the golang structure of table device for DAO operations like Where/Data.
type Device struct {
	g.Meta     `orm:"table:device, do:true"`
	Id         interface{} //
	DeviceId   interface{} // 自定义设备ID
	Mac        interface{} // mac地址
	Ip         interface{} // Ip地址
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	ActiveTime *gtime.Time // 最后活跃时间
	Able       interface{} // 是否启用
	DeviceName interface{} // 设备名字
	DeviceImg  interface{} // 设备图片
	Remark     interface{} //
	UserId     interface{} // 绑定的用户
}
