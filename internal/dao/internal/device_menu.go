// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// DeviceMenuDao is the data access object for the table device_menu.
type DeviceMenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  DeviceMenuColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// DeviceMenuColumns defines and stores column names for the table device_menu.
type DeviceMenuColumns struct {
	Id         string // 主键
	Level      string // 层级
	Pid        string // 父ID
	Code       string // 编码
	Name       string // 名字
	CreateTime string // 创建时间
	UpdateTime string // 更新时间
	Type       string // 类型：0文件夹，1路由页面，3接口
	Path       string // 路径
	Icon       string // 图标
	Order      string // 排序
	Select     string // 选择
}

// deviceMenuColumns holds the columns for the table device_menu.
var deviceMenuColumns = DeviceMenuColumns{
	Id:         "id",
	Level:      "level",
	Pid:        "pid",
	Code:       "code",
	Name:       "name",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Type:       "type",
	Path:       "path",
	Icon:       "icon",
	Order:      "order",
	Select:     "select",
}

// NewDeviceMenuDao creates and returns a new DAO object for table data access.
func NewDeviceMenuDao(handlers ...gdb.ModelHandler) *DeviceMenuDao {
	return &DeviceMenuDao{
		group:    "default",
		table:    "device_menu",
		columns:  deviceMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *DeviceMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *DeviceMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *DeviceMenuDao) Columns() DeviceMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *DeviceMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *DeviceMenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *DeviceMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
