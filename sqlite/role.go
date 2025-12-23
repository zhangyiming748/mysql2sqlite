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