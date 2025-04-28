// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// License is the golang structure of table license for DAO operations like Where/Data.
type License struct {
	g.Meta       `orm:"table:license, do:true"`
	Id           interface{} //
	LicenseId    interface{} // 许可证ID
	DeviceId     interface{} // 设备ID
	DeviceName   interface{} // 设备名字
	StartTime    *gtime.Time // 开始时间
	EndTime      *gtime.Time // 结束时间
	CreateTime   *gtime.Time // 创建时间
	UpdateTime   *gtime.Time // 更新时间
	Type         interface{} // 许可证订阅类型
	PayImg       interface{} // 支付截图
	PayType      interface{} // 支付类型
	Money        interface{} // 金钱
	Remark       interface{} // 备注
	SerialNumber interface{} // 流水号
	Able         interface{} // 是否启用状态
}
