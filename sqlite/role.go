package sqlite

import (
	"time"
)

type Role struct {
	Id          int       `gorm:"primaryKey;autoIncrement;type:INTEGER;comment:主键id"`
	RoleName    string    `gorm:"type:TEXT;comment:角色"`
	Description string    `gorm:"type:TEXT;comment:简短角色描述"`
	ParentRole  string    `gorm:"type:TEXT;comment:根角色"`
	CreatedAt   time.Time `gorm:"type:DATETIME"`
	UpdatedAt   time.Time `gorm:"type:DATETIME"`
	DeletedAt   time.Time `gorm:"type:DATETIME"`
}

/*
获取表中全部记录到结构体切片
*/
func (r *Role) GetAll() (roles []Role, err error) {
	err = GetSqlite().Find(&roles).Error
	return roles, err
}

/*
同步表结构到数据库
*/
func (r *Role) SyncRole() error {
	db := GetSqlite()
	return db.AutoMigrate(&Role{})
}

/*
添加单条记录到数据库
*/
func (r *Role) InsertOne() error {
	db := GetSqlite()
	return db.Create(r).Error
}