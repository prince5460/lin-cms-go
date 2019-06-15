package core

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var Config = beego.AppConfig

func InitEnv() {
	var err error
	// database
	dbUser := Config.String("mysqluser")
	dbPass := Config.String("mysqlpass")
	dbHost := Config.String("mysqlhost")
	dbPort := Config.String("mysqlport")
	dbName := Config.String("mysqldb")
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName) + "&loc=Asia%2FChongqing"
	if err = orm.RegisterDriver("mysql", orm.DRMySQL); err != nil {
		panic(err)
	}
	if err = orm.RegisterDataBase("default", "mysql", dbLink); err != nil {
		panic(err)
	}
	if Config.String("runmode") == "dev" {
		orm.Debug = true
	}

}
