package license

import (
	"context"
	v1 "gf-admin/api/license/v1"
	"gf-admin/internal/dao"
)

func (l *License) List(ctx context.Context, req *v1.ListReq) (*v1.ListRes, error) {
	var (
		result = &v1.ListRes{}
		items  []v1.License
	)

	// 构建查询
	m := dao.License.Ctx(ctx).As("l")

	// LEFT JOIN 关联查询
	m = m.LeftJoin("device d", "l.device_id=d.device_id").
		LeftJoin("device_user u", "d.user_id=u.id").
		Fields(`
			l.*,
			d.device_name as DeviceName,
			u.user_name as UserName,
			u.id as UserId
		`)

	// 分页查询
	m.Page(req.PageNo, req.PageSize).ScanAndCount(&items, &result.Total, false)

	result.PageNo = req.PageNo
	result.PageSize = req.PageSize
	result.Records = items

	return result, nil
}
