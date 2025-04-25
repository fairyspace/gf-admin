// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceMenu is the golang structure of table device_menu for DAO operations like Where/Data.
type DeviceMenu struct {
	g.Meta     `orm:"table:device_menu, do:true"`
	Id         interface{} // 主键
	Level      interface{} // 层级
	Pid        interface{} // 父ID
	Code       interface{} // 编码
	Name       interface{} // 名字
	CreateTime *gtime.Time // 创建时间
	UpdateTime *gtime.Time // 更新时间
	Type       interface{} // 类型：0文件夹，1路由页面，3接口
	Path       interface{} // 路径
	Icon       interface{} // 图标
	Order      interface{} // 排序
	Select     interface{} // 选择
}
