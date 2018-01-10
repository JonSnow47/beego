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

func (a *AdminController) Login() {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug(admin)
		a.Data["json"] = "Request error"
	} else {
		err := models.AdminServer.Create(admin)
		if err != nil {
			a.Data["json"] = "New false"
		} else {
			info, _ := models.AdminServer.AdminInfo(admin.Name)
			a.Data["json"] = map[string]interface{}{"info": info}
		}
	}
	a.ServeJSON()
}

func (a *AdminController) Signin() {
	var Logininfo struct {
		name     string
		password string
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &Logininfo)
	if err != nil {
		logs.Debug(Logininfo)
		a.Data["json"] = logs.Debug
	} else {
		err := models.AdminServer.Signin(Logininfo.name, Logininfo.password)
		if err != nil {
			logs.Debug("Name or password error.")
			a.Data["json"] = "SQL error."
		} else {
			a.Data["json"] = "Sign in."
		}
	}
	a.ServeJSON()
}

/*
func (a *AdminController) Signout() {

}
*/
