package device_menu

import (
	"context"
	v1 "gf-admin/api/device_menu/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (d *DeviceMenu) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	orm := dao.DeviceMenu.Ctx(ctx)

	// 检查菜单非自身Code是否重复
	count, err := orm.Where(do.DeviceMenu{
		Code: req.Code,
	}).WhereNot("id", req.Id).Count()

	if err != nil {
		return nil, err
	}

	if count > 0 {
		return nil, gerror.New("菜单编码已存在")
	}

	// 更新菜单
	_, err = orm.Data(do.DeviceMenu{
		Id:         req.Id,
		Code:       req.Code,
		Name:       req.Name,
		Path:       req.Path,
		Icon:       req.Icon,
		Order:      req.Order,
		Type:       req.Type,
		Level:      req.Level,
		Pid:        req.Pid,
		Select:     req.Select,
		UpdateTime: gtime.Now(),
	}).Where("id", req.Id).Update()

	if err != nil {
		return nil, err
	}

	return &v1.UpdateRes{}, nil
}
