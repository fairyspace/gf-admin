// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package device_role

import (
	api "gf-admin/api/device_role"
	logic "gf-admin/internal/logic/device_role"
)

type ControllerV1 struct {
	deviceRole *logic.DeviceRole
}

func NewV1() api.IDeviceRoleV1 {
	return &ControllerV1{
		deviceRole: logic.New(),
	}
}
