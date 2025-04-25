package device_role

import (
	"context"

	v1 "gf-admin/api/device_role/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return c.deviceRole.Update(ctx, req)
}
