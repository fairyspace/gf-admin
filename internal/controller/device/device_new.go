// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package device

import (
	"gf-admin/api/device"
	logic "gf-admin/internal/logic/device"
)

type ControllerV1 struct {
	device *logic.Device
}

func NewV1() device.IDeviceV1 {
	return &ControllerV1{
		device: logic.New(),
	}
}
