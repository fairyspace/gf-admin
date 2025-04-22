// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceUser is the golang structure of table device_user for DAO operations like Where/Data.
type DeviceUser struct {
	g.Meta     `orm:"table:device_user, do:true"`
	Id         interface{} // 主键
	UserName   interface{} // 用户
	Password   interface{} // 密码
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Avatar     interface{} // 头像
	Active     interface{} // 是否禁用
}
