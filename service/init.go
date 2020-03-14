package service

import (
	"GO_IM/model"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var DbEngin *xorm.Engine

func init() {
	driverName := "mysql"
	DsName := "root:root@(127.0.0.1:3306)/goim?charset=utf8"
	var err error
	DbEngin, err = xorm.NewEngine(driverName, DsName)
	if err != nil {
		log.Fatal(err.Error())
	}
	//显示sql语句
	DbEngin.ShowSQL(true)

	//设置数据库最大连接数
	DbEngin.SetMaxOpenConns(2)

	//自动同步表结构
	DbEngin.Sync2(new(model.User),
		new(model.Contact),
		new(model.Community))

	fmt.Println("Init data ok.")
}
