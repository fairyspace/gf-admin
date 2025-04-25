package device_role

import (
	"context"
	v1 "gf-admin/api/device_role/v1"
	"gf-admin/internal/dao"
)

func (d *DeviceRole) Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error) {
	dao.DeviceRole.Ctx(ctx).WhereIn("id", req.Ids).Delete()
	return &v1.DeleteRes{}, nil
}
