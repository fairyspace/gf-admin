package device

import (
	"context"
	v1 "gf-admin/api/device/v1"
	"gf-admin/internal/dao"
	"gf-admin/internal/model/entity"
)

type TempUser struct {
	Id       int
	UserName string
}

func (d *Device) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {

	res = &v1.GetListRes{} // 修复：初始化res，防止nil指针
	//1.当userName不为空时，模糊查询device_user表中userName对应的userId
	var TempUser []TempUser
	var userIds []int
	userNameMap := make(map[int]string)

	if req.UserName != "" {
		orm := dao.DeviceUser.Ctx(ctx)
		orm.Where("user_name LIKE ?", "%"+req.UserName+"%").Scan(&TempUser)
		// 提取userIds切片
		for _, user := range TempUser {
			userIds = append(userIds, user.Id)
		}

		//建立userId-userName的Map映射

		for _, user := range TempUser {
			userNameMap[user.Id] = user.UserName
		}
	}

	// 2.当deviceId不为空时，模糊查询deviceId对应的设备列表
	var devices []*entity.Device
	orm := dao.Device.Ctx(ctx)

	if len(userIds) > 0 {
		orm = orm.WhereIn("user_id", userIds)
	}

	if req.DeviceId != "" {
		orm = orm.WhereLike("device_id", "%"+req.DeviceId+"%")
	}
	//分页
	orm.Page(req.PageNo, req.PageSize).ScanAndCount(&devices, &res.Total, false)
	// 3.将devices转换为v1.Device
	var records []v1.Device
	for _, device := range devices {
		records = append(records, v1.Device{
			Id:         device.Id,
			DeviceId:   device.DeviceId,
			DeviceName: device.DeviceName,
			Mac:        device.Mac,
			Ip:         device.Ip,
			DeviceImg:  device.DeviceImg,
			Remark:     device.Remark,
			Able:       device.Able == 1,
			UserId:     device.UserId,
			UserName:   userNameMap[device.UserId],
		})
	}
	res.Records = records

	res.PageNo = req.PageNo
	res.PageSize = req.PageSize

	return

}
