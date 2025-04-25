// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceRoleDao is the data access object for the table device_role.
type DeviceRoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DeviceRoleColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DeviceRoleColumns defines and stores column names for the table device_role.
type DeviceRoleColumns struct {
	Id         string // 主键
	RoleName   string // 角色名称
	CreateTime string //
	UpdateTime string //
	Remark     string //
}

// deviceRoleColumns holds the columns for the table device_role.
var deviceRoleColumns = DeviceRoleColumns{
	Id:         "id",
	RoleName:   "role_name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Remark:     "remark",
}

// NewDeviceRoleDao creates and returns a new DAO object for table data access.
func NewDeviceRoleDao(handlers ...gdb.ModelHandler) *DeviceRoleDao {
	return &DeviceRoleDao{
		group:    "default",
		table:    "device_role",
		columns:  deviceRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DeviceRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DeviceRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DeviceRoleDao) Columns() DeviceRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DeviceRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DeviceRoleDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *DeviceRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
