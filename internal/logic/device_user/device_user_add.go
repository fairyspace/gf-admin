package device_user

import (
	"context"
	v1 "gf-admin/api/device_user/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (u *DeviceUser) Add(ctx context.Context, req *v1.AddReq) (*v1.AddRes, error) {
	// 调用dao层新增设备用户
	orm := dao.DeviceUser.Ctx(ctx)

	// 查询是否存在相同用户名
	count, err := orm.Where(do.DeviceUser{
		UserName: req.UserName,
	}).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("用户名已存在")
	}

	// 新增设备用户
	_, err = orm.Data(do.DeviceUser{
		UserName:   req.UserName,
		Avatar:     req.Avatar,
		Password:   req.Password,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	}).Insert()

	if err != nil {
		return nil, err
	}

	return &v1.AddRes{}, nil
}
