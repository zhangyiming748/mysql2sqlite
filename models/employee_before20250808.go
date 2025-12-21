package models

import (
	"time"
)

type EmployeeBefore20250808 struct {
	Id                int       `xorm:"not null pk autoincr comment('主键id') INT"`
	EmployId          string    `xorm:"comment('员工编号') VARCHAR(128)"`
	EmployName        string    `xorm:"comment('员工姓名') VARCHAR(128)"`
	EmployRole        string    `xorm:"comment('用户角色') VARCHAR(128)"`
	EmployDescription string    `xorm:"comment('简短角色描述') VARCHAR(128)"`
	EmployGroup       string    `xorm:"VARCHAR(255)"`
	CreatedAt         time.Time `xorm:"DATETIME"`
	UpdatedAt         time.Time `xorm:"DATETIME"`
	DeletedAt         time.Time `xorm:"DATETIME"`
}
