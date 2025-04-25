// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceMenu is the golang structure for table device_menu.
type DeviceMenu struct {
	Id         int         `json:"id"         orm:"id"          description:"主键"`                // 主键
	Level      int         `json:"level"      orm:"level"       description:"层级"`                // 层级
	Pid        int         `json:"pid"        orm:"pid"         description:"父ID"`               // 父ID
	Code       string      `json:"code"       orm:"code"        description:"编码"`                // 编码
	Name       string      `json:"name"       orm:"name"        description:"名字"`                // 名字
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"`              // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"`              // 更新时间
	Type       int         `json:"type"       orm:"type"        description:"类型：0文件夹，1路由页面，3接口"` // 类型：0文件夹，1路由页面，3接口
	Path       string      `json:"path"       orm:"path"        description:"路径"`                // 路径
	Icon       string      `json:"icon"       orm:"icon"        description:"图标"`                // 图标
	Order      int         `json:"order"      orm:"order"       description:"排序"`                // 排序
	Select     int         `json:"select"     orm:"select"      description:"选择"`                // 选择
}
