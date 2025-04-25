package device_menu

import (
	"context"
	v1 "gf-admin/api/device_menu/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"
	"gf-admin/utility"
)

func (d *DeviceMenu) GetList(ctx context.Context, req *v1.GetListReq) (*v1.GetListRes, error) {
	// 获取所有菜单记录
	var menuList []*entity.DeviceMenu
	err := dao.DeviceMenu.Ctx(ctx).
		Scan(&menuList)
	if err != nil {
		return nil, err
	}

	// 构建菜单树
	menuTree := utility.BuildMenuTree(menuList, 0)

	return &menuTree, nil
}
