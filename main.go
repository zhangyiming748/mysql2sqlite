package main

import (
	"log"
	"mysql2sqlite/mysql"
	"mysql2sqlite/sqlite"
	"os"
)

func initConfig() {
	mc := &mysql.MysqlConfig{
		DbName:   "Permissions",
		Host:     "192.168.110.68",
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
	new(sqlite.Position).SyncPosition()
	new(sqlite.Role).SyncRole()
}

func main() {
	initConfig()
	if roles, err := new(mysql.Role).GetAll(); err != nil {
		log.Fatalf("%v\n", err)
	} else {
		for i, role := range roles {
			log.Printf("%d:%+v\n", i+1, role)
			r:=new(sqlite.Role)
			r.RoleName = role.RoleName
			r.Description = role.Description
			if err =r.InsertOne();err!=nil{
				log.Fatalf("%v\n", err)			}
		}
	}
	if positions, err := new(mysql.Position).GetAll(); err != nil {
		log.Fatalf("%v\n", err)
	} else {
		for i, position := range positions {
			log.Printf("%d:%+v\n", i+1, position)
			p:=new(sqlite.Position)
			p.PositionNumber = position.PositionNumber
			p.PositionName = position.PositionName
			p.ParentRole = position.ParentRole
			p.RoleDesc = position.RoleDesc
			p.RoleIdentifier = position.RoleIdentifier
			p.MenuType = position.MenuType
			p.MenuNode = position.MenuNode
			p.TCodeDesc = position.TCodeDesc
			p.IsCoreTCode = position.IsCoreTCode
			p.NeedDocument = position.NeedDocument
			p.StorageTable = position.StorageTable
			if err = p.InsertOne(); err != nil {
				log.Fatalf("%v\n", err)
			}
		}
	}
}