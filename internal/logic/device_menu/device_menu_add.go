package device_menu

import (
	"context"
	v1 "gf-admin/api/device_menu/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"
	"gf-admin/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (d *DeviceMenu) Add(ctx context.Context, req *v1.AddReq) (*v1.AddRes, error) {
	// 获取orm
	orm := dao.DeviceMenu.Ctx(ctx)

	//检查菜单是否重复,除去自身
	var menu *entity.DeviceMenu
	err := orm.Where(do.DeviceMenu{
		Code: req.Code,
	}).Scan(&menu)

	if err != nil {
		return nil, err
	}
	// 检查菜单是否重复
	if menu != nil {
		return nil, gerror.New("菜单已存在")
	}

	// 添加菜单
	_, err = orm.Data(do.DeviceMenu{
		Code:       req.Code,
		Name:       req.Name,
		Path:       req.Path,
		Icon:       req.Icon,
		Order:      req.Order,
		Type:       req.Type,
		Level:      req.Level,
		Pid:        req.Pid,
		Select:     req.Select,
		CreateTime: gtime.Now(),
		UpdateTime: gtime.Now(),
	}).Insert()

	if err != nil {
		return nil, err
	}

	return &v1.AddRes{}, nil
}
