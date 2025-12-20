package mysql

import (
	"log"
	"testing"
)

func init() {
	mc := &MysqlConfig{
		DbName:   "tdl",
		Host:     "192.168.5.2",
		Password: "163453",
		Port:     "3306",
		User:     "root",
	}
	if err := SetMysql(mc); err != nil {
		log.Fatal(err)
	}
	log.SetFlags(log.Ltime | log.Lshortfile)
}

// go test -v -run TestGetAll
func TestGetAll(t *testing.T) {
	var files []File
	files = new(File).GetAllFile()
	for _, file := range files {
		t.Log(file.Filename)
		// t.Logf("%v\n",file)
	}
}
