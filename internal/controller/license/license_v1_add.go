package license

import (
	"context"

	"gf-admin/api/license/v1"
)

func (c *ControllerV1) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	return c.license.Add(ctx, req)
}
