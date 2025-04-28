package device

import (
	"context"
	v1 "gf-admin/api/device/v1"
	"gf-admin/internal/dao"
)

func (d *Device) Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error) {
	// 获取orm
	orm := dao.Device.Ctx(ctx)
	// 删除
	_, err := orm.Where("id IN (?)", req.Ids).Delete()
	if err != nil {
		return nil, err
	}
	return &v1.DeleteRes{}, nil
}
