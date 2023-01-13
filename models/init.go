package models

import (
	_ "database/sql"
	"fmt"

	"time"
	// _ "ginstudy/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

//读数据
var orm *xorm.Engine

// //写数据
// var DB_Write *xorm.Engine

// var engine *xorm.Engine
var err error

func init() {
	orm, err = xorm.NewEngine("mysql", "root:root@tcp(localhost:3306)/ginstudy?charset=utf8")
	// engine, err := xorm.NewEngine("mysql", "2343432:122222@/(http://127.0.0.1:3306)/ginstudy?charset=utf8")
	// db, err = xorm.NewEngine("mysql", "username:password@tcp(host:3306)/dbname?charset=utf8")
	if err != nil {
		fmt.Println(err)
		fmt.Println(err.Error())
		return
	} else {
		orm.Ping()
		//是否显示sql语句
		orm.ShowSQL(true)
		if err = orm.Sync2(new(Admin), new(User), new(Authgroup), new(Authrule), new(Authaccess), new(City), new(Category), new(News), new(Device)); err != nil {
			fmt.Println(err)
		} else {
			// orm.ShowWarn(true)
			// engine.ShowWarn(true)
			fmt.Print("自动生成表成功！")
		}

	}
	orm.SetMaxOpenConns(50)
	// 设置连接池的空闲数大小
	orm.SetMaxIdleConns(5)
	// 设置空闲连接最大时长
	// engine.SetConnMaxIdleTime(600) xorm没有这个功能
	// 设置连接最大存活时长，必须小于mysql的wait_timeout
	// mysql wait_timeout单位是s，time.Duration 默认是纳秒，后面*1000000000，转化成s
	// 设置为4h
	orm.SetConnMaxLifetime(59 * time.Second)
	// fmt.Println(err)
}
