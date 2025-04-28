package license

import (
	"context"

	"gf-admin/api/license/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	return c.license.Delete(ctx, req)
}
