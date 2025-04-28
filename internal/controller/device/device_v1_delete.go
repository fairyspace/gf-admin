package device

import (
	"context"

	v1 "gf-admin/api/device/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return c.device.Delete(ctx, req)
}
