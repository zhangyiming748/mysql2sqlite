package models

import (
	"time"
)

type EmployeeBefore366 struct {
	Id                int       `xorm:"not null pk autoincr comment('主键id') INT"`
	EmployId          string    `xorm:"comment('用户ID') VARCHAR(512)"`
	EmployName        string    `xorm:"comment('用户姓名') VARCHAR(512)"`
	EmployRole        string    `xorm:"comment('用户角色') VARCHAR(512)"`
	EmployGroup       string    `xorm:"VARCHAR(255)"`
	EmployDescription string    `xorm:"comment('简短角色描述') VARCHAR(512)"`
	CreatedAt         time.Time `xorm:"DATETIME"`
	UpdatedAt         time.Time `xorm:"DATETIME"`
	DeletedAt         time.Time `xorm:"DATETIME"`
}
