package main

import (
	"github.com/JonSnow47/beego/application/Fan/filters"
	"github.com/JonSnow47/beego/application/travel/initialize"
	_ "github.com/JonSnow47/beego/application/travel/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
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
