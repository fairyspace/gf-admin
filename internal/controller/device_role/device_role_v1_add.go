package device_role

import (
	"context"

	v1 "gf-admin/api/device_role/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	return c.deviceRole.Add(ctx, req)
}
