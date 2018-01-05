package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/travel/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
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
	info := models.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	if err != nil {
		logs.Debug(info)
		this.Data["json"] = "error"
	} else {
		models.UserServer.Register(info)
		this.Data["json"] = map[string]interface{}{"result": "Successful",
			"id":   info.ID,
			"name": info.Name}
	}
	this.ServeJSON()
}

func (this *UserController) UserInfo() {
	//info := models.User{}
	var info models.User
	id := info.ID
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		this.Data["json"] = "error"
	} else {
		info := models.UserServer.ReadInfobyID(id)
		this.Data["json"] = info
	}
	//this.Data["json"] = map[string]string{"content": "43"}
	this.ServeJSON()
}
/*
func (this *UserController) LogOff() {
	id := models.User.ID
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		this.Data["json"] = "error"

	}
	objectid := models.AddOne(ob)
	this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"
	this.ServeJSON()
}
*/