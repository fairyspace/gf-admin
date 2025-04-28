package license

import (
	"context"
	"gf-admin/api/license/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/do"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

func (l License) Add(ctx context.Context, req *v1.AddReq) (*v1.AddRes, error) {
	orml := dao.License.Ctx(ctx)

	//检查licenseId是否重复
	count, err := orml.Where(do.License{LicenseId: req.LicenseId}).Count()
	if err != nil {
		return nil, gerror.Wrap(err, "查询licenseId重复失败")
	}

	if count > 0 {
		return nil, gerror.New("licenseId已存在")
	}

	//插入数据
	_, err = orml.Data(do.License{
		LicenseId:    req.LicenseId,
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
		CreateTime:   gtime.Now(),
		UpdateTime:   gtime.Now(),
	}).Insert()

	if err != nil {
		return nil, err
	}

	return &v1.AddRes{}, err

}
