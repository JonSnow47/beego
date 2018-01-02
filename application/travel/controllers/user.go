package controllers

import (
	"github.com/JonSnow47/beego/application/travel/models"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Hello() {
	this.Data["json"] = models.Hello("Jon Snow")
	this.ServeJSON()
}

func (this *UserController) Register() {
	Nickname := this.GetString("nickname")
	if Nickname == "" {
		this.Ctx.WriteString("Please choose a nickname.")
		return
	} else if len(Nickname) > 16 {
		this.Ctx.WriteString("Please choose a shorter nickname.")
	} else {
		this.Ctx.WriteString(Nickname)
	}
}

func (this *UserController) UserInfo() {
	this.Data["json"] = map[string]string{"content": "43"}
}
