package device

import (
	"context"

	v1 "gf-admin/api/device/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return c.device.Update(ctx, req)
}
