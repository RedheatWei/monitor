package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/astaxie/beego"
	"fmt"
)

func ConnDB() *xorm.Engine{
	dbconn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%utf8",beego.AppConfig.String("mysqluser"),beego.AppConfig.String("mysqlpass"),beego.AppConfig.String("mysqlhost"),beego.AppConfig.String("mysqldb"))
	var err error
	engine, err = xorm.NewEngine("mysql", dbconn)
	if err!=nil{
		fmt.Println(err)
	}
	return engine
}