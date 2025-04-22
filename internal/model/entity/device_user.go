// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceUser is the golang structure for table device_user.
type DeviceUser struct {
	Id         int         `json:"id"         orm:"id"          description:"主键"`   // 主键
	UserName   string      `json:"userName"   orm:"user_name"   description:"用户"`   // 用户
	Password   string      `json:"password"   orm:"password"    description:"密码"`   // 密码
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:"创建时间"` // 创建时间
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:"更新时间"` // 更新时间
	Avatar     string      `json:"avatar"     orm:"avatar"      description:"头像"`   // 头像
	Active     int         `json:"active"     orm:"active"      description:"是否禁用"` // 是否禁用
}
