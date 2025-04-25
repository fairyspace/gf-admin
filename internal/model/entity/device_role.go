// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// DeviceRole is the golang structure for table device_role.
type DeviceRole struct {
	Id         int         `json:"id"         orm:"id"          description:"主键"`   // 主键
	RoleName   string      `json:"roleName"   orm:"role_name"   description:"角色名称"` // 角色名称
	CreateTime *gtime.Time `json:"createTime" orm:"create_time" description:""`     //
	UpdateTime *gtime.Time `json:"updateTime" orm:"update_time" description:""`     //
	Remark     string      `json:"remark"     orm:"remark"      description:""`     //
}
