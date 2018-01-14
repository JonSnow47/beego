package main

import (
	_ "github.com/JonSnow47/beego/application/travel/routers"
	"github.com/JonSnow47/beego/application/travel/initialize"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	initialize.InitMysql()
	beego.Run()
}
