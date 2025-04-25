package device_role

import (
	"context"
	v1 "gf-admin/api/device_role/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
)

func (d *DeviceRole) GetList(ctx context.Context, req *v1.GetListReq) (*v1.GetListRes, error) {
	var list []*entity.DeviceRole
	var res v1.GetListRes

	orm := dao.DeviceRole.Ctx(ctx)

	// 分页查询
	if req.RoleName != "" {
		orm.Where(do.DeviceRole{
			RoleName: req.RoleName,
		})
	}

	orm.Page(req.PageNo, req.PageSize).ScanAndCount(&list, &res.Total, false)

	// 设置分页信息
	res.PageNo = req.PageNo
	res.PageSize = req.PageSize

	// 设置查询结果
	res.Records = list

	return &res, nil
}
