package device_role

import (
	"context"

	v1 "gf-admin/api/device_role/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {

	return c.deviceRole.GetList(ctx, req)
}
