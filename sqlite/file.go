package sqlite

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type File struct {
	Id      int64  `gorm:"primaryKey;autoIncrement;comment:主键id"`
	Origin  string `gorm:"column:origin;type:varchar(255);comment:来源"`
	Channel string `gorm:"column:channel;type:varchar(100);comment:频道"`
	Fid     int    `gorm:"column:fid;type:int;comment:文件ID"`
	Tag     string `gorm:"column:tag;type:varchar(100);comment:标签"`
	Subtag  string `gorm:"column:subtag;type:varchar(100);comment:子标签"`
	Fname   string `gorm:"column:fname;type:varchar(255);comment:文件名"`
	source  string `gorm:"column:source;type:varchar(255);comment:源"`

	CreatedAt time.Time      // GORM 会自动管理这些时间字段
	UpdatedAt time.Time      // GORM 会自动管理这些时间字段
	DeletedAt gorm.DeletedAt `gorm:"index"` // 软删除支持
}

// 同步表结构
func SyncFile() {
	// 自动迁移表结构
	err := GetSqlite().AutoMigrate(&File{})
	if err != nil {
		log.Fatalf("同步表结构失败: %v", err)
	}

	log.Println("File表结构同步成功")
}

func (f *File) InsertOne() (success bool, err error) {
	//这里实现插入一条数据
	db := GetSqlite()
	if db == nil {
		return false, fmt.Errorf("数据库连接未初始化")
	}

	result := db.Create(f)
	if result.Error != nil {
		return false, result.Error
	}

	return true, nil
}
