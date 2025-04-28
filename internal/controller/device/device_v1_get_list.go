package device

import (
	"context"

	v1 "gf-admin/api/device/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	return c.device.GetList(ctx, req)
}
