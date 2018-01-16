package main

import (
	"github.com/astaxie/beego"

	"github.com/astaxie/beego/plugins/cors"

	"github.com/JonSnow47/beego/application/Fan/initialize"
	"github.com/JonSnow47/beego/application/Fan/filters"
	_ "github.com/JonSnow47/beego/application/Fan/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}))
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LoginFilter)
	initialize.InitMysql()
	beego.Run()
}
