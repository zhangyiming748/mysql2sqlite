package models

import (
	"time"
)

type MmRolePurchasePlan struct {
	Id                         int       `xorm:"not null pk autoincr comment('主键id') INT"`
	ErpFunctionPositionID      string    `xorm:"comment('ERP功能岗位序号') VARCHAR(512)"`
	ErpFunctionPosition        string    `xorm:"comment('ERP功能岗位') VARCHAR(512)"`
	RootRole                   string    `xorm:"comment('根角色') VARCHAR(512)"`
	RootRoleDescription        string    `xorm:"comment('根角色描述') VARCHAR(512)"`
	TransactionCode            string    `xorm:"comment('事务代码') VARCHAR(512)"`
	TransactionCodeDescription string    `xorm:"comment('事务代码或其他（菜单节点）描述') VARCHAR(512)"`
	CreatedAt                  time.Time `xorm:"DATETIME"`
	UpdatedAt                  time.Time `xorm:"DATETIME"`
	DeletedAt                  time.Time `xorm:"DATETIME"`
}
