package device_user

import (
	"context"
	v1 "gf-admin/api/device_user/v1"
	"gf-admin/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (u *DeviceUser) Set(ctx context.Context, req *v1.SetReq) (*v1.SetRes, error) {
	//开启事务
	err := dao.DeviceUser.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		//删除原有的角色
		_, err := tx.Model("ass_user_role").Where("user_id", req.Id).Delete()
		if err != nil {
			return err
		}

		//批量添加新的角色
		batchData := make([]g.Map, 0, len(req.RoleIds))
		for _, roleId := range req.RoleIds {
			batchData = append(batchData, g.Map{
				"user_id":     req.Id,
				"role_id":     roleId,
				"create_time": gtime.Now(),
				"update_time": gtime.Now(),
			})
		}
		_, err = tx.Model("ass_user_role").Data(batchData).Batch(len(batchData)).Insert()
		if err != nil {
			return err
		}

		return nil
	})

	return &v1.SetRes{}, err
}
