// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceRole is the golang structure of table device_role for DAO operations like Where/Data.
type DeviceRole struct {
	g.Meta     `orm:"table:device_role, do:true"`
	Id         interface{} // 主键
	RoleName   interface{} // 角色名称
	CreateTime *gtime.Time //
	UpdateTime *gtime.Time //
	Remark     interface{} //
}
