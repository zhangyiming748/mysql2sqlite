package models

import (
	"time"
)

type EmployeeIdentifierInfo struct {
	Id         int       `xorm:"not null pk autoincr comment('主键id') INT"`
	EmployeeID string    `xorm:"comment('职工编码') VARCHAR(512)"`
	Name       string    `xorm:"comment('姓名') VARCHAR(512)"`
	CnpcEmail  string    `xorm:"comment('中石油邮箱') VARCHAR(512)"`
	Mobile     string    `xorm:"comment('手机号') VARCHAR(512)"`
	CreatedAt  time.Time `xorm:"DATETIME"`
	UpdatedAt  time.Time `xorm:"DATETIME"`
	DeletedAt  time.Time `xorm:"DATETIME"`
}
