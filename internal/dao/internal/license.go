// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// LicenseDao is the data access object for the table license.
type LicenseDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  LicenseColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// LicenseColumns defines and stores column names for the table license.
type LicenseColumns struct {
	Id           string //
	LicenseId    string // 许可证ID
	DeviceId     string // 设备ID
	DeviceName   string // 设备名字
	StartTime    string // 开始时间
	EndTime      string // 结束时间
	CreateTime   string // 创建时间
	UpdateTime   string // 更新时间
	Type         string // 许可证订阅类型
	PayImg       string // 支付截图
	PayType      string // 支付类型
	Money        string // 金钱
	Remark       string // 备注
	SerialNumber string // 流水号
	Able         string // 是否启用状态
}

// licenseColumns holds the columns for the table license.
var licenseColumns = LicenseColumns{
	Id:           "id",
	LicenseId:    "license_id",
	DeviceId:     "device_id",
	DeviceName:   "device_name",
	StartTime:    "start_time",
	EndTime:      "end_time",
	CreateTime:   "create_time",
	UpdateTime:   "update_time",
	Type:         "type",
	PayImg:       "pay_img",
	PayType:      "pay_type",
	Money:        "money",
	Remark:       "remark",
	SerialNumber: "serial_number",
	Able:         "able",
}

// NewLicenseDao creates and returns a new DAO object for table data access.
func NewLicenseDao(handlers ...gdb.ModelHandler) *LicenseDao {
	return &LicenseDao{
		group:    "default",
		table:    "license",
		columns:  licenseColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *LicenseDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *LicenseDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *LicenseDao) Columns() LicenseColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *LicenseDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *LicenseDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *LicenseDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
