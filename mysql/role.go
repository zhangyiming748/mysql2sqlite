package mysql

import (
	"time"
)

type Role struct {
	Id          int       `xorm:"not null pk autoincr comment('主键id') INT"`
	RoleName    string    `xorm:"comment('角色') VARCHAR(512)"`
	Description string    `xorm:"comment('简短角色描述') VARCHAR(512)"`
	ParentRole  string    `xorm:"comment('根角色') VARCHAR(512)"`
	CreatedAt   time.Time `xorm:"DATETIME"`
	UpdatedAt   time.Time `xorm:"DATETIME"`
	DeletedAt   time.Time `xorm:"DATETIME"`
}
