package controllers

import (
	"github.com/JonSnow47/beego/application/travel/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (this *UserController) Hello() {

	this.Data["json"] = models.Hello("Jon Snow")
	this.ServeJSON()
}

func (this *UserController) UserID() {
	this.Data["json"] = models.ID(47)
	this.ServeJSON()
}