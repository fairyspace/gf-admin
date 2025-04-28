package license

import (
	"context"
	"gf-admin/api/license/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"
)

func (l License) Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error) {
	orm := dao.License.Ctx(ctx)
	orm.Where(do.License{Id: req.Id}).Delete()

	return &v1.DeleteRes{}, nil
}
