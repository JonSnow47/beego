package main

import (
	"github.com/astaxie/beego"

	"github.com/JonSnow47/beego/application/Fan/initialize"
	_ "github.com/JonSnow47/beego/application/Fan/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	initialize.InitMysql()
	beego.Run()
}
