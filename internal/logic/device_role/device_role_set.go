package device_role

import (
	"context"
	v1 "gf-admin/api/device_role/v1"
	"gf-admin/internal/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

func (d *DeviceRole) Set(ctx context.Context, req *v1.SetReq) (*v1.SetRes, error) {
	g.Log().Info(ctx, "req", req)

	// 开启事务
	err := dao.DeviceRole.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 清除旧的菜单数据
		_, err := tx.Model("ass_role_menu").Where("role_id", req.Id).Delete()
		if err != nil {
			return err
		}

		// 添加新的菜单数据
		if len(req.MenuIds) > 0 {
			// 构建批量插入的数据
			batchData := make([]g.Map, 0, len(req.MenuIds))
			for _, menuId := range req.MenuIds {
				batchData = append(batchData, g.Map{
					"role_id":     req.Id,
					"menu_id":     menuId,
					"create_time": gtime.Now(),
					"update_time": gtime.Now(),
				})
			}

			// 执行批量插入
			_, err := tx.Model("ass_role_menu").Data(batchData).Batch(len(batchData)).Insert()
			if err != nil {
				return err
			}
		}

		return nil
	})

	return &v1.SetRes{}, err
}
