package models

import (
	"time"
)

type MmRole struct {
	Id                     int       `xorm:"not null pk autoincr comment('主键id') INT"`
	RootRole               string    `xorm:"comment('根角色') VARCHAR(512)"`
	RootRoleDescription    string    `xorm:"comment('根角色描述') VARCHAR(512)"`
	DerivedRole            string    `xorm:"comment('派生角色') VARCHAR(512)"`
	DerivedRoleDescription string    `xorm:"comment('派生角色描述') VARCHAR(512)"`
	CreatedAt              time.Time `xorm:"DATETIME"`
	UpdatedAt              time.Time `xorm:"DATETIME"`
	DeletedAt              time.Time `xorm:"DATETIME"`
}
