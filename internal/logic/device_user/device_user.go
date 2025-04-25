package device_user

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

type DeviceUser struct {
}

func New() *DeviceUser {
	return &DeviceUser{}
}

// GetUserRoleIds 获取用户角色ID列表
func GetUserRoleIds(ctx context.Context, userId int) ([]int, error) {
	values, err := g.DB().Model("ass_user_role").Where("user_id", userId).Array("role_id")
	if err != nil {
		return nil, err
	}

	var roleIds []int
	for _, value := range values {
		roleIds = append(roleIds, value.Int())
	}

	return roleIds, nil
}
