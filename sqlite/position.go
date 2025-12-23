package sqlite

import (
	"time"
)

type Position struct {
	Id             int       `gorm:"primaryKey;autoIncrement;type:INTEGER;comment:主键id"`
	PositionNumber string    `gorm:"type:TEXT;comment:ERP功能岗位序号"`
	PositionName   string    `gorm:"type:TEXT;comment:ERP功能岗位"`
	ParentRole     string    `gorm:"type:TEXT;comment:根角色"`
	RoleDesc       string    `gorm:"type:TEXT;comment:角色描述"`
	RoleIdentifier string    `gorm:"type:TEXT;comment:根角色标识"`
	MenuType       string    `gorm:"type:TEXT;comment:菜单节点类型"`
	MenuNode       string    `gorm:"type:TEXT;comment:菜单节点"`
	TCodeDesc      string    `gorm:"type:TEXT;comment:事务代码或其他(菜单节点)描述"`
	IsCoreTCode    int       `gorm:"type:INTEGER;comment:是否核心TCODE"`
	NeedDocument   int       `gorm:"type:INTEGER;comment:是否要求产生单据"`
	StorageTable   string    `gorm:"type:TEXT;comment:产生单据后台主要存储表"`
	CreatedAt      time.Time `gorm:"type:DATETIME"`
	UpdatedAt      time.Time `gorm:"type:DATETIME"`
	DeletedAt      time.Time `gorm:"type:DATETIME"`
}

/*
获取表中全部记录到结构体切片
*/
func (p *Position) GetAll() (positions []Position, err error) {
	err = GetSqlite().Find(&positions).Error
	return positions, err
}

/*
同步表结构到数据库
*/
func (p *Position) SyncPosition() error {
	db := GetSqlite()
	return db.AutoMigrate(&Position{})
}

/*
添加单条记录到数据库
*/
func (p *Position) InsertOne() error {
	db := GetSqlite()
	return db.Create(p).Error
}