package license

import (
	"context"
	v1 "gf-admin/api/license/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"
	"github.com/gogf/gf/v2/os/gtime"
)

func (l License) Update(ctx context.Context, req *v1.UpdateReq) (*v1.UpdateRes, error) {
	orm := dao.License.Ctx(ctx)
	//根据ID更新数据
	// 执行更新操作
	orm.Where(do.License{Id: req.Id}).Data(do.License{LicenseId: req.LicenseId,
		DeviceId:     req.DeviceId,
		DeviceName:   req.DeviceName,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Type:         req.Type,
		PayImg:       req.PayImg,
		PayType:      req.PayType,
		Money:        req.Money,
		Remark:       req.Remark,
		SerialNumber: req.SerialNumber,
		Able:         req.Able,
		UpdateTime:   gtime.Now(),
	}).Update()

	return &v1.UpdateRes{}, nil
}
