package license

import (
	"context"

	"gf-admin/api/license/v1"
)

func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	return c.license.Update(ctx, req)
}
