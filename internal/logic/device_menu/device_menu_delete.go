package device_menu

import (
	"context"
	v1 "gf-admin/api/device_menu/v1"
	"gf-admin/internal/dao"
)

func (d *DeviceMenu) Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error) {
	orm := dao.DeviceMenu.Ctx(ctx)

	_, err := orm.WhereIn("id", req.Ids).Delete()
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRes{}, nil
}
