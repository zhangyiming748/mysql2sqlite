package models

import (
	"time"
)

type BpmApprovalNode struct {
	Id             int       `xorm:"not null pk autoincr comment('主键id') INT"`
	ApprovalPolicy string    `xorm:"comment('审批策略') VARCHAR(512)"`
	OrgCode        string    `xorm:"comment('组织编码') VARCHAR(512)"`
	OrgName        string    `xorm:"comment('组织名称') VARCHAR(512)"`
	RoleCode       string    `xorm:"comment('角色编码') VARCHAR(512)"`
	RoleName       string    `xorm:"comment('角色名称') VARCHAR(512)"`
	ApprovalType   string    `xorm:"comment('审批类型') VARCHAR(512)"`
	ApprovalLevel  string    `xorm:"comment('审批级别') VARCHAR(512)"`
	UserAccount    string    `xorm:"comment('用户账号') VARCHAR(512)"`
	UserName       string    `xorm:"comment('用户名称') VARCHAR(512)"`
	CreatedAt      time.Time `xorm:"DATETIME"`
	UpdatedAt      time.Time `xorm:"DATETIME"`
	DeletedAt      time.Time `xorm:"DATETIME"`
}
