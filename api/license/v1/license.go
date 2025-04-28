package v1

import (
	"gf-admin/internal/consts"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type AddReq struct {
	g.Meta       `path:"/add" method:"post" summary:"Add a license"`
	LicenseId    int                `json:"licenseId"`
	DeviceId     string             `json:"deviceId"`
	DeviceName   string             `json:"deviceName"`
	StartTime    *gtime.Time        `json:"startTime"`
	EndTime      *gtime.Time        `json:"endTime"`
	Type         consts.LicenseType `json:"type"` // 使用公共枚举类型
	PayImg       string             `json:"payImg"`
	PayType      string             `json:"payType"`
	Money        float64            `json:"money"`
	Remark       string             `json:"remark"`
	SerialNumber string             `json:"serialNumber"`
	Able         bool               `json:"able"`
}

type AddRes struct {
}

type UpdateReq struct {
	g.Meta       `path:"/update" method:"put" summary:"Update a license"`
	Id           int                `json:"id"`
	LicenseId    int                `json:"licenseId"`
	DeviceId     string             `json:"deviceId"`
	DeviceName   string             `json:"deviceName"`
	StartTime    *gtime.Time        `json:"startTime"`
	EndTime      *gtime.Time        `json:"endTime"`
	Type         consts.LicenseType `json:"type"` // 使用公共枚举类型
	PayImg       string             `json:"payImg"`
	PayType      string             `json:"payType"`
	Money        float64            `json:"money"`
	Remark       string             `json:"remark"`
	SerialNumber string             `json:"serialNumber"`
	Able         bool               `json:"able"`
}

type UpdateRes struct{}

type DeleteReq struct {
	g.Meta `path:"/delete" method:"delete" summary:"Delete a license"`
	Id     int `json:"id"`
}

type DeleteRes struct{}

type ListReq struct {
	g.Meta   `path:"/list" method:"post" summary:"List licenses"`
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
}

type ListRes struct {
	Total    int       `json:"total"`
	PageNo   int       `json:"pageNo"`
	PageSize int       `json:"pageSize"`
	Records  []License `json:"records"`
}

type License struct {
	Id           int         `json:"id"`
	LicenseId    string      `json:"licenseId"`
	DeviceId     string      `json:"deviceId"`
	DeviceName   string      `json:"deviceName"`
	StartTime    *gtime.Time `json:"startTime"`
	EndTime      *gtime.Time `json:"endTime"`
	CreateTime   *gtime.Time `json:"createTime"`
	UpdateTime   *gtime.Time `json:"updateTime"`
	Type         string      `json:"type"`
	PayImg       string      `json:"payImg"`
	PayType      string      `json:"payType"`
	Money        float64     `json:"money"`
	Remark       string      `json:"remark"`
	SerialNumber string      `json:"serialNumber"`
	Able         int         `json:"able"`
	UserId       int         `json:"userId"`
	UserName     string      `json:"userName"`
}
