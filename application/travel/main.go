package main

import (
	_ "github.com/JonSnow47/beego/application/travel/routers"
	"github.com/JonSnow47/beego/application/travel/initorm"
	"github.com/astaxie/beego"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	initorm.InitMysql()
	beego.Run()
}
