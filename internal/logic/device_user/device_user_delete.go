package device_user

import (
	"context"
	v1 "gf-admin/api/device_user/v1"
	"gf-admin/internal/dao"
)

func (u *DeviceUser) Delete(ctx context.Context, req *v1.DeleteReq) (*v1.DeleteRes, error) {
	// 调用dao层删除设备用户
	orm := dao.DeviceUser.Ctx(ctx)

	// 批量删除设备用户
	_, err := orm.WhereIn("id", req.Ids).Delete()
	if err != nil {
		return nil, err
	}

	return &v1.DeleteRes{}, nil
}
