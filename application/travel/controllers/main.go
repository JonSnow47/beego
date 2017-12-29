package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func(this *MainController) Get()  {
	this.Data["json"] = "Welcome to Travel!"
	this.ServeJSON()
}