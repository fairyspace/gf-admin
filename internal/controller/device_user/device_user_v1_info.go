package device_user

import (
	"context"

	v1 "gf-admin/api/device_user/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	return c.deviceUser.Info(ctx, req)
}
