package device

import (
	"context"

	v1 "gf-admin/api/device/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	return c.device.Add(ctx, req)
}
