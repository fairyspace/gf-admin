package device_menu

import (
	"context"

	v1 "gf-admin/api/device_menu/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return c.deviceMenu.Delete(ctx, req)
}
