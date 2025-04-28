package device

import (
	"context"
	v1 "gf-admin/api/device/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (d *Device) Add(ctx context.Context, req *v1.AddReq) (*v1.AddRes, error) {
	//获取ORM
	orm := dao.Device.Ctx(ctx)

	//判断IP是否重复
	count, err := orm.Where("ip = ?", req.Ip).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("IP已存在")
	}

	//判断deviceId是否重复
	count, err = orm.Where("device_id = ?", req.DeviceId).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("deviceId已存在")
	}
	// req转entity
	orm.Data(do.Device{
		DeviceId:   req.DeviceId,
		DeviceName: req.DeviceName,
		Mac:        req.Mac,
		Ip:         req.Ip,
		DeviceImg:  req.DeviceImg,
		Remark:     req.Remark,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
		Able:       req.Able,
	}).Insert()

	return &v1.AddRes{}, nil
}
