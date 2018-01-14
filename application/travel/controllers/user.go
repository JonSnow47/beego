package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/JonSnow47/beego/application/travel/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Hello() {
	this.Data["json"] = models.Hello("Jon Snow")
	this.ServeJSON()
}

func (this *UserController) Register() {
	info := models.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &info)
	if err != nil {
		logs.Debug(info)
		this.Data["json"] = "Request error"
	} else {
		models.UserServer.Register(info)
		this.Data["json"] = map[string]interface{}{"result": "Register Success",
			"id":   info.ID,
			"name": info.Name}
	}
	this.ServeJSON()
}

func (this *UserController) UserInfo() {
	var Id struct{
		id	int64
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &Id)
	if err != nil {
		logs.Debug(Id)
		this.Data["json"] = "Request error"
	} else {
		content, err := models.UserServer.ReadInfoByID(Id.id)
		this.Data["json"] = content
		if err != nil {
			logs.Debug(content)
			this.Data["json"] = "error"
		} else {
			this.Data["json"] = content
		}
	}
	//this.Data["json"] = map[string]string{"content": "43"}
	this.ServeJSON()
}

func (this *UserController) Delete() {
	content := models.User{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &content)
	if err != nil {
		logs.Debug(content)
		this.Data["json"] = "error"
	} else {
		models.UserServer.Delete(content)
		this.Data["json"] = map[string]interface{}{"result":"delete success"}
	}
	/*objectid := models.AddOne(ob)
	this.Data["json"] = "{\"ObjectId\":\"" + objectid + "\"}"*/
	this.ServeJSON()
}