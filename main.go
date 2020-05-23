package main

import (
	_ "employmentInfo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "root:daomei,815@tcp(127.0.0.1:3306)/employmentinfo?charset=utf8")
	beego.Run()
}
