package license

import (
	"context"

	v1 "gf-admin/api/license/v1"
)

func (c *ControllerV1) List(ctx context.Context, req *v1.ListReq) (res *v1.ListRes, err error) {
	return c.license.List(ctx, req)
}
