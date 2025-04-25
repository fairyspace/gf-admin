package device_role

import (
	"context"
	v1 "gf-admin/api/device_role/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"
	"gf-admin/utility"

	"github.com/gogf/gf/v2/frame/g"
)

func (d *DeviceRole) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoRes, error) {

	// 2. 查询关联的菜单ID列表

	values, _ := g.DB().Model("ass_role_menu").Where("role_id", req.Id).Array("menu_id")
	var menuIds []int
	for _, id := range values {
		menuIds = append(menuIds, id.Int())
	}

	// 3. Inquire About Menu Details
	var menuList []*entity.DeviceMenu
	dao.DeviceMenu.Ctx(ctx).WhereIn("id", menuIds).Scan(&menuList)

	menuTree := utility.BuildMenuTree(menuList, 0)

	return &menuTree, nil
}
