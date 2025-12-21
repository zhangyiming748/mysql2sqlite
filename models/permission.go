package models

import (
	"time"
)

type Permission struct {
	Id         int       `xorm:"not null pk autoincr comment('主键id') INT"`
	EmployeeID string    `xorm:"comment('职工编码') VARCHAR(512)"`
	Name       string    `xorm:"comment('姓名') VARCHAR(512)"`
	Role       string    `xorm:"comment('角色') VARCHAR(512)"`
	RoleDesc   string    `xorm:"comment('简短角色描述') VARCHAR(512)"`
	CreatedAt  time.Time `xorm:"DATETIME"`
	UpdatedAt  time.Time `xorm:"DATETIME"`
	DeletedAt  time.Time `xorm:"DATETIME"`
}
