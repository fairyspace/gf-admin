package device_user

import (
	"context"

	v1 "gf-admin/api/device_user/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return c.deviceUser.Update(ctx, req)
}
