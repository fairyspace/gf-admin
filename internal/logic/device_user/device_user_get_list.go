package device_user

import (
	"context"
	v1 "gf-admin/api/device_user/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"
)

// GetList 获取设备用户列表
// 支持分页查询和用户名模糊搜索
func (u *DeviceUser) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	// 初始化返回结构
	res = &v1.GetListRes{}
	// 初始化列表切片
	list := make([]*entity.DeviceUser, 0)

	// 构建查询条件
	orm := dao.DeviceUser.Ctx(ctx)

	// 只有当用户名不为空时才添加用户名过滤条件
	if req.UserName != "" {
		orm = orm.Where(do.DeviceUser{
			UserName: req.UserName,
		})
	}

	// 执行分页查询
	orm.Page(req.PageNo, req.PageSize).ScanAndCount(&list, &res.Total, false)

	// 设置分页信息
	res.PageNo = req.PageNo
	res.PageSize = req.PageSize
	// 设置查询结果
	res.Records = list

	return res, nil
}
