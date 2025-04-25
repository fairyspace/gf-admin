package device_user

import (
	"context"
	v1 "gf-admin/api/device_user/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

func (u *DeviceUser) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoRes, error) {
	//查询用户-角色关系表ass_user_role

	values, _ := g.DB().Model("ass_user_role").Where("user_id", req.Id).Array("user_id")

	var roleIds []int
	for _, value := range values {
		roleIds = append(roleIds, value.Int())
	}

	//查询角色列表
	var roleList []*entity.DeviceRole
	dao.DeviceRole.Ctx(ctx).WhereIn("id", roleIds).Scan(&roleList)

	//查询所有角色列表
	var allRoleList []*entity.DeviceRole
	dao.DeviceRole.Ctx(ctx).Scan(&allRoleList)

	return &v1.InfoRes{
		AssignRoles: roleList,
		AllRoleList: allRoleList,
	}, nil
}
