package device_role

import (
	"context"
	v1 "gf-admin/api/device_role/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (d *DeviceRole) Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error) {
	orm := dao.DeviceRole.Ctx(ctx)

	// 检查角色是否存在
	count, err := orm.Where(do.DeviceRole{
		RoleName: req.RoleName,
	}).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("角色已存在")
	}

	// 新增角色
	_, err = orm.Data(do.DeviceRole{
		RoleName:   req.RoleName,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	}).Insert()

	if err != nil {
		return nil, err
	}

	return &v1.AddRes{}, nil
}
