package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

/*
这个函数用来实现xorm框架的mysql连接初始化
第一次连接不要选择数据库
连接后判断是否存在给定的数据库
如果存在直接连接
否则创建数据库
然后再连接
*/
type MysqlConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

var Engine *xorm.Engine

func SetMysql(mc *MysqlConfig) error {
	// 创建不指定数据库的连接字符串
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/", mc.User, mc.Password, mc.Host, mc.Port)

	// 第一次连接不要选择数据库
	tempEngine, err := xorm.NewEngine("mysql", dsnWithoutDB)
	if err != nil {
		return err
	}
	defer tempEngine.Close()

	// 检查数据库是否存在，通过执行SQL查询
	exists := false
	rows, err := tempEngine.DB().Query("SHOW DATABASES LIKE '" + mc.DbName + "'")
	if err != nil {
		return err
	}

	if rows.Next() {
		exists = true
	}
	rows.Close()

	// 如果不存在则创建数据库
	if !exists {
		_, err = tempEngine.Exec("CREATE DATABASE " + mc.DbName)
		if err != nil {
			return err
		}
	}

	// 创建指向目标数据库的连接字符串
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mc.User, mc.Password, mc.Host, mc.Port, mc.DbName)

	// 创建持久化引擎
	Engine, err = xorm.NewEngine("mysql", dsnWithDB)
	if err != nil {
		return err
	}

	return nil
}

func GetEngine() *xorm.Engine {
	return Engine
}
