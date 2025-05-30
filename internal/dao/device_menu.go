// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"gf-admin/internal/dao/internal"
)

// deviceMenuDao is the data access object for the table device_menu.
// You can define custom methods on it to extend its functionality as needed.
type deviceMenuDao struct {
	*internal.DeviceMenuDao
}

var (
	// DeviceMenu is a globally accessible object for table device_menu operations.
	DeviceMenu = deviceMenuDao{internal.NewDeviceMenuDao()}
)

// Add your custom methods and functionality below.
