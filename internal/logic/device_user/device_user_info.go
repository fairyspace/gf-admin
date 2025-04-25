package device_user

import (
	"context"
	v1 "gf-admin/api/device_user/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"
)

func (u *DeviceUser) Info(ctx context.Context, req *v1.InfoReq) (*v1.InfoRes, error) {
	//查询用户-角色关系表ass_user_role
	roleIds, err := GetUserRoleIds(ctx, req.Id)
	if err != nil {
		return nil, err
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
