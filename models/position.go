package models

import (
	"time"
)

type Position struct {
	Id             int       `xorm:"not null pk autoincr comment('主键id') INT"`
	PositionNumber string    `xorm:"comment('ERP功能岗位序号') VARCHAR(512)"`
	PositionName   string    `xorm:"comment('ERP功能岗位') VARCHAR(512)"`
	ParentRole     string    `xorm:"comment('根角色') VARCHAR(512)"`
	RoleDesc       string    `xorm:"comment('角色描述') VARCHAR(512)"`
	RoleIdentifier string    `xorm:"comment('根角色标识') VARCHAR(512)"`
	MenuType       string    `xorm:"comment('菜单节点类型') VARCHAR(512)"`
	MenuNode       string    `xorm:"comment('菜单节点') VARCHAR(512)"`
	TCodeDesc      string    `xorm:"comment('事务代码或其他(菜单节点)描述') VARCHAR(512)"`
	IsCoreTCode    int       `xorm:"comment('是否核心TCODE') TINYINT(1)"`
	NeedDocument   int       `xorm:"comment('是否要求产生单据') TINYINT(1)"`
	StorageTable   string    `xorm:"comment('产生单据后台主要存储表') VARCHAR(512)"`
	CreatedAt      time.Time `xorm:"DATETIME"`
	UpdatedAt      time.Time `xorm:"DATETIME"`
	DeletedAt      time.Time `xorm:"DATETIME"`
}
