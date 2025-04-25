package login

import (
	"context"

	v1 "gf-admin/api/login/v1"
)

func (c *ControllerV1) Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error) {
	return c.login.Info(ctx, req)
}
