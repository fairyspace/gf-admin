package login

import (
	"context"

	v1 "gf-admin/api/login/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	return c.login.Logout(ctx, req)
}
