package models

import (
	"time"
)

type MmTransactionRole struct {
	Id                             int       `xorm:"not null pk autoincr comment('主键id') INT"`
	ApprovalCode                   string    `xorm:"VARCHAR(255)"`
	ApprovalCodeDescription        string    `xorm:"comment('审批代码描述') VARCHAR(512)"`
	RootApprovalRole               string    `xorm:"comment('审批根角色') VARCHAR(512)"`
	DerivedApprovalRole            string    `xorm:"comment('审批派生角色:') VARCHAR(512)"`
	DerivedApprovalRoleDescription string    `xorm:"comment('审批派生角色描述') VARCHAR(512)"`
	CreatedAt                      time.Time `xorm:"DATETIME"`
	UpdatedAt                      time.Time `xorm:"DATETIME"`
	DeletedAt                      time.Time `xorm:"DATETIME"`
}
