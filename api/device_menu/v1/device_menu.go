package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetListReq struct {
	g.Meta `path:"/list" method:"post"`
}

type Menu struct {
	Id       int    `json:"id"`
	Pid      int    `json:"pid"`
	Level    int    `json:"level"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Icon     string `json:"icon"`
	Order    int    `json:"order"`
	Type     int    `json:"type"`
	Select   bool   `json:"select"`
	Children []Menu `json:"children"`
}

// 返回数组
type GetListRes = []Menu

type AddReq struct {
	g.Meta `path:"/add" method:"post"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Icon   string `json:"icon"`
	Order  int    `json:"order"`
	Type   int    `json:"type"`
	Level  int    `json:"level"`
	Pid    int    `json:"pid"`
	Select bool   `json:"select"`
}

type AddRes struct {
}

type UpdateReq struct {
	g.Meta `path:"/update" method:"put"`
	Id     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Icon   string `json:"icon"`
	Order  int    `json:"order"`
	Type   int    `json:"type"`
	Level  int    `json:"level"`
	Pid    int    `json:"pid"`
	Select bool   `json:"select"`
}

type UpdateRes struct {
}

type DeleteReq struct {
	g.Meta `path:"/delete" method:"delete"`
	Ids    []int `json:"ids"`
}

type DeleteRes struct {
}
