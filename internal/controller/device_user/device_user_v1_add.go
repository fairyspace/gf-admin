package device_user

import (
	"context"

	v1 "gf-admin/api/device_user/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	// 调用逻辑层新增设备用户
	return c.deviceUser.Add(ctx, req)
}
