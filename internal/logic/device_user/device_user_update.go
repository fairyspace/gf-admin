package device_user

import (
	"context"
	v1 "gf-admin/api/device_user/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
)

// Update 更新设备用户信息
func (u *DeviceUser) Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error) {
	// 调用dao层更新设备用户信息
	//获取ORM
	orm := dao.DeviceUser.Ctx(ctx)

	// 查询是否存在相同用户名,但是排除自己
	count, err := orm.Where(do.DeviceUser{
		UserName: req.UserName,
	}).WhereNot("id", req.ID).Count()

	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, gerror.New("用户名已存在")
	}

	// 更新设备用户信息
	_, err = orm.Where(do.DeviceUser{
		Id: req.ID,
	}).Data(do.DeviceUser{
		UserName: req.UserName,
	}).Update()

	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{}, nil
}
