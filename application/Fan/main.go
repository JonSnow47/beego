package main

import (
	_ "github.com/JonSnow47/beego/application/Fan/routers"
	"github.com/JonSnow47/beego/application/Fan/orm"

	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.InitMysql()
	beego.Run()
}
