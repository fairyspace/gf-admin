// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceUserDao is the data access object for the table device_user.
type DeviceUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DeviceUserColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DeviceUserColumns defines and stores column names for the table device_user.
type DeviceUserColumns struct {
	Id         string // 主键
	UserName   string // 用户
	Password   string // 密码
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Avatar     string // 头像
	Active     string // 是否禁用
}

// deviceUserColumns holds the columns for the table device_user.
var deviceUserColumns = DeviceUserColumns{
	Id:         "id",
	UserName:   "user_name",
	Password:   "password",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Avatar:     "avatar",
	Active:     "active",
}

// NewDeviceUserDao creates and returns a new DAO object for table data access.
func NewDeviceUserDao(handlers ...gdb.ModelHandler) *DeviceUserDao {
	return &DeviceUserDao{
		group:    "default",
		table:    "device_user",
		columns:  deviceUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DeviceUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DeviceUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DeviceUserDao) Columns() DeviceUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DeviceUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DeviceUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DeviceUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
