package sqlite

import (
	"log"
	"path/filepath"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func SetSqlite(root string) {
	location := filepath.Join(root, "sqlite.db")
	db, err := gorm.Open(sqlite.Open(location), &gorm.Config{})
	if err != nil {
		log.Fatalf("打开本地sqlite数据库失败:%s", err.Error())
	}

	gormDB = db
	log.Println("本地sqlite数据库初始化完成")
}

func GetSqlite() *gorm.DB {
	return gormDB
}
