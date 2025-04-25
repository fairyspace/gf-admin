package device_menu

import (
	"context"

	v1 "gf-admin/api/device_menu/v1"
)

func (c *ControllerV1) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	return c.deviceMenu.GetList(ctx, req)
}
