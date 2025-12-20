package main

import (
	"log"
	"mysql2sqlite/mysql"
	"mysql2sqlite/sqlite"
	"os"
)

func initConfig() {
	mc := &mysql.MysqlConfig{
		DbName:   "tdl",
		Host:     "192.168.5.2",
		Password: "163453",
		Port:     "3306",
		User:     "root",
	}
	if err := mysql.SetMysql(mc); err != nil {
		log.Fatal(err)
	}
	log.SetFlags(log.Ltime | log.Lshortfile)
	if home, err := os.UserHomeDir(); err != nil {
		log.Fatal(err)
	} else {
		sqlite.SetSqlite(home)
	}
}

func main() {
	initConfig()
	var files []mysql.File
	files = new(mysql.File).GetAllFile()
	for _, file := range files {
		log.Println(file.Filename)
		s := new(sqlite.File)
		s.Fname = file.Filename
		s.InsertOne()
	}
}
