package device_role

import (
	"context"
	v1 "gf-admin/api/device_role/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (d *DeviceRole) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	orm := dao.DeviceRole.Ctx(ctx)

	// 检查角色是否存在,排除自己
	count, err := orm.Where(do.DeviceRole{
		Id: req.Id,
	}).WhereNot("id", req.Id).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("角色已存在")
	}

	// 更新角色
	_, err = orm.Data(do.DeviceRole{
		RoleName: req.RoleName,
	}).Update()
	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{}, nil
}
