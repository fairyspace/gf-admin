package login

import (
	"context"
	v1 "gf-admin/api/login/v1"
	"gf-admin/internal/consts"
	"gf-admin/internal/dao"
	"gf-admin/internal/logic/device_user"
	"gf-admin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (l *Login) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoRes, error) {
	// 从上下文中获取用户ID
	userId := ctx.Value(consts.CtxUserId).(uint)

	// 查询用户信息
	var user *entity.DeviceUser
	err := dao.DeviceUser.Ctx(ctx).Where("id", userId).Scan(&user)
	if err != nil {
		return nil, err
	}

	// 查询用户有哪些角色 ass_user_role
	roleIds, err := device_user.GetUserRoleIds(ctx, int(userId))
	if err != nil {
		return nil, err
	}

	// 查询角色信息
	var roles []string
	roleNameValues, err := dao.DeviceRole.Ctx(ctx).WhereIn("id", roleIds).Array("role_name")

	for _, value := range roleNameValues {
		roles = append(roles, value.String())
	}

	if err != nil {
		return nil, err
	}

	// 查询角色对应的菜单Ids,操作关系表ass_role_menu
	values, _ := g.DB().Model("ass_role_menu").WhereIn("role_id", roleIds).Array("menu_id")
	var menuIds []int
	for _, id := range values {
		menuIds = append(menuIds, id.Int())
	}

	// 查询菜单信息
	var menus []*entity.DeviceMenu
	err = dao.DeviceMenu.Ctx(ctx).WhereIn("id", menuIds).Scan(&menus)
	if err != nil {
		return nil, err
	}

	// menus根据type=1和type=2分组，再映射成code切片
	var routeList []string
	var buttonList []string
	for _, menu := range menus {
		if menu.Type == 1 {
			routeList = append(routeList, menu.Code)
		}
		if menu.Type == 2 {
			buttonList = append(buttonList, menu.Code)
		}
	}

	return &v1.InfoRes{
		UserId:   userId,
		Avatar:   user.Avatar,
		UserName: user.UserName,
		Buttons:  buttonList,
		Roles:    roles,
		Routes:   routeList,
	}, nil
}
