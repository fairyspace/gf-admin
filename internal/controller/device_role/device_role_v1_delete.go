package device_role

import (
	"context"

	v1 "gf-admin/api/device_role/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return c.deviceRole.Delete(ctx, req)
}
