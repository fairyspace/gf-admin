package v1

import (
	"gf-admin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"

	menu "gf-admin/api/device_menu/v1"
)

type GetListReq struct {
	g.Meta   `path:"/list" method:"post"`
	PageNo   int    `p:"pageNo" d:"1" v:"min:1#页码不能小于1"`
	PageSize int    `p:"pageSize" d:"10" v:"min:1#每页数量不能小于1"`
	RoleName string `p:"roleName" json:"roleName"`
}

type GetListRes struct {
	Total    int                  `json:"total"`
	PageNo   int                  `json:"pageNo"`
	PageSize int                  `json:"pageSize"`
	Records  []*entity.DeviceRole `json:"records"`
}

// 新增角色
type AddReq struct {
	g.Meta   `path:"/add" method:"post"`
	RoleName string `p:"roleName" json:"roleName"`
}

type AddRes struct {
}

type UpdateReq struct {
	g.Meta   `path:"/update" method:"put"`
	RoleName string `p:"roleName" json:"roleName"`
	Id       int    `p:"id" json:"id"`
}

type UpdateRes struct {
}

// 删除角色
type DeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Ids    []int `p:"ids" json:"ids"`
}

type DeleteRes struct {
}

// 获取角色菜单
type InfoReq struct {
	g.Meta `path:"/info" method:"get"`
	Id     int `p:"id" json:"id"`
}

type InfoRes = []menu.Menu

// 设置角色菜单
type SetReq struct {
	g.Meta  `path:"/set" method:"post"`
	Id      int   `p:"id" json:"id"`
	MenuIds []int `p:"menuIds" json:"menuIds"`
}

type SetRes struct {
}
