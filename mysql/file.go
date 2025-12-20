package mysql

import (
	"log"
	"time"
)

type File struct {
	Id        int64     `xorm:"not null pk autoincr INT(11)"`
	Origin    string    `xorm:"VARCHAR(255)"`
	Channel   string    `xorm:"VARCHAR(255)"`
	FileId    int       `xorm:"INT(11)"`
	Tag       string    `xorm:"VARCHAR(255)"`
	Subtag    string    `xorm:"VARCHAR(255)"`
	Filename  string    `xorm:"VARCHAR(255)"`
	Offset    int       `xorm:"INT(11)"`
	Capacity  int       `xorm:"INT(11)"`
	From      string    `xorm:"VARCHAR(255)"`
	CreatedAt time.Time `xorm:"DATETIME"`
	UpdatedAt time.Time `xorm:"DATETIME"`
	DeletedAt time.Time `xorm:"DATETIME"`
}

func (f *File) Sync() {
	log.Printf("开始同步表结构")
	// 注意：原 storage.GetMysql().Sync2 调用已被移除，因依赖无效
	log.Printf("同步表结构功能暂不可用，请实现有效的数据库同步逻辑")
}

// 这个函数用来返回全部数据的结构体切片
func (f *File) GetAllFile() []File {
	var file []File
	GetEngine().Find(&file)
	return file
}
