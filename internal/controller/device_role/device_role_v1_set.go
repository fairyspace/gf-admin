package device_role

import (
	"context"

	v1 "gf-admin/api/device_role/v1"
)

func (c *ControllerV1) Set(ctx context.Context, req *v1.SetReq) (res *v1.SetRes, err error) {
	return c.deviceRole.Set(ctx, req)
}
