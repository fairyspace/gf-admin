package device_user

import (
	"context"

	v1 "gf-admin/api/device_user/v1"
)

func (c *ControllerV1) Set(ctx context.Context, req *v1.SetReq) (res *v1.SetRes, err error) {
	return c.deviceUser.Set(ctx, req)
}
