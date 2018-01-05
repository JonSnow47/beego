package controllers

import (
	"encoding/json"
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
/*
func (this *UserController) Register() {

	nickname := this.GetString("nickname")
	if nickname == "" {
		this.Ctx.WriteString("Please choose a nickname.")
		return
	} else if len(nickname) > 16 {
		this.Ctx.WriteString("Please choose a shorter nickname.")
	} else {
		this.Ctx.WriteString(nickname)
	}
}
*/
func (this *UserController) Register() {
	var info models.User
	json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	models.UserServer.Register(&info)
	this.ServeJSON()
}

func (this *UserController) UserInfo() {
	this.Data["json"] = map[string]string{"content": "43"}
}
