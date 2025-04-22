package v1

import (
	"gf-admin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// 获取设备用户列表请求
type GetListReq struct {
	g.Meta   `path:"/list" method:"post" tags:"设备用户" summary:"获取设备用户列表"`
	PageNo   int    `json:"pageNo" d:"1" v:"min:1#页码不能小于1"`
	PageSize int    `json:"pageSize" d:"10" v:"min:1#每页数量不能小于1"`
	UserName string `json:"userName"`
}

// 获取设备用户列表响应
type GetListRes struct {
	g.Meta   `mime:"application/json"`
	Total    int                  `json:"total"`
	PageNo   int                  `json:"pageNo"`
	PageSize int                  `json:"pageSize"`
	Records  []*entity.DeviceUser `json:"records"`
}

// 修改设备用户信息请求
type UpdateReq struct {
	g.Meta   `path:"/update" method:"put" tags:"设备用户" summary:"修改设备用户信息"`
	ID       int    `json:"id" v:"required#ID不能为空"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Active   int    `json:"active"`
}

// 修改设备用户信息响应
type UpdateRes struct {
	g.Meta `mime:"application/json"`
}

// 新增设备用户请求
type AddReq struct {
	g.Meta   `path:"/add" method:"post" tags:"设备用户" summary:"新增设备用户"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
}

// 新增设备用户响应
type AddRes struct {
	g.Meta `mime:"application/json"`
}

// 删除设备用户请求
type DeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"设备用户" summary:"删除设备用户"`
	Ids    []int `json:"ids" v:"required#Ids不能为空"`
}

// 删除设备用户响应
type DeleteRes struct {
	g.Meta `mime:"application/json"`
}
