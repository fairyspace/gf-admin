// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package device_role

import (
	"context"

	"gf-admin/api/device_role/v1"
)

type IDeviceRoleV1 interface {
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	Add(ctx context.Context, req *v1.AddReq) (res *v1.AddRes, err error)
	Update(ctx context.Context, req *v1.UpdateReq) (res *v1.UpdateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	Info(ctx context.Context, req *v1.InfoReq) (res *v1.InfoRes, err error)
	Set(ctx context.Context, req *v1.SetReq) (res *v1.SetRes, err error)
}
