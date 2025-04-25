package login

import (
	"context"

	v1 "gf-admin/api/login/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return c.login.Login(ctx, req)
}
