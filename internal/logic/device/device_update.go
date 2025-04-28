package device

import (
	"context"
	v1 "gf-admin/api/device/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (d *Device) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	orm := dao.Device.Ctx(ctx)

	//判断deviceId是否存在
	count, err := orm.Where("device_id = ?", req.DeviceId).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("设备不存在")
	}

	//更新设备
	orm.Where("id = ?", req.Id).Data(do.Device{
		Id:         req.Id,
		DeviceId:   req.DeviceId,
		DeviceName: req.DeviceName,
		Mac:        req.Mac,
		Ip:         req.Ip,
		DeviceImg:  req.DeviceImg,
		Remark:     req.Remark,
		Able:       req.Able,
	}).Update()

	return &v1.UpdateRes{}, nil
}
