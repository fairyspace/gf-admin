package v1

import "github.com/gogf/gf/v2/frame/g"

type AddReq struct {
	g.Meta     `path:"/add" method:"post" tags:"设备" summary:"添加设备"`
	DeviceId   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	Mac        string `json:"mac"`
	Ip         string `json:"ip"`
	DeviceImg  string `json:"deviceImg"`
	Remark     string `json:"remark"`
	Able       bool   `json:"able"`
	UserId     int    `json:"userId"`
}

type AddRes struct {
}

type UpdateReq struct {
	g.Meta     `path:"/update" method:"put" tags:"设备" summary:"更新设备"`
	DeviceId   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	Mac        string `json:"mac"`
	Ip         string `json:"ip"`
	DeviceImg  string `json:"deviceImg"`
	Remark     string `json:"remark"`
	Able       bool   `json:"able"`
	Id         int    `json:"id"`
	UserId     int    `json:"userId"`
}

type UpdateRes struct {
}

type DeleteReq struct {
	g.Meta `path:"/delete" method:"delete" tags:"设备" summary:"删除设备"`
	Ids    []int `json:"ids"`
}

type DeleteRes struct {
}

type GetListReq struct {
	g.Meta        `path:"/list" method:"post" tags:"设备" summary:"获取设备列表"`
	PageNo        int    `json:"pageNo"`
	PageSize      int    `json:"pageSize"`
	DeviceId      string `json:"deviceId"`
	UserName      string `json:"userName"`
	LicenseStatus string `json:"licenseStatus"`
}

type GetListRes struct {
	Total    int      `json:"total"`
	PageNo   int      `json:"pageNo"`
	PageSize int      `json:"pageSize"`
	Records  []Device `json:"records"`
}

type Device struct {
	Id         int    `json:"id"`
	DeviceId   string `json:"deviceId"`
	DeviceName string `json:"deviceName"`
	Mac        string `json:"mac"`
	Ip         string `json:"ip"`
	DeviceImg  string `json:"deviceImg"`
	Remark     string `json:"remark"`
	Able       bool   `json:"able"`
	UserId     int    `json:"userId"`
	UserName   string `json:"userName"`
}
