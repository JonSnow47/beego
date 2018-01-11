package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/Fan/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AdminController struct {
	beego.Controller
}

func (a *AdminController) New() {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug(admin)
		a.Data["json"] = "Request error"
	} else {
		err := models.AdminServer.Create(admin.Name, admin.Password)
		if err != nil {
			a.Data["json"] = "New false"
		} else {
			a.Data["json"] = map[string]interface{}{"New admin": admin.Name}
		}
	}
	a.ServeJSON()
}

func (a *AdminController) Login() {
	var Loginfo struct {
		name     string `json:"name"`
		password string `json:"password"`
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &Loginfo)
	if err != nil {
		logs.Debug(Loginfo)
		a.Data["json"] = "Request error."
	} else {
		ok, err := models.AdminServer.Login(Loginfo.name, Loginfo.password)
		if err != nil {
			logs.Debug(err)
			a.Data["json"] = "Error."
		} else if err == nil && ok == true {
			a.Data["json"] = map[string]interface{}{"Log in": Loginfo.name}
		} else {
			a.Data["json"] = "Incorrect username or password."
		}
	}
	a.ServeJSON()
}

/*
func (a *AdminController) Logout() {

}
*/
