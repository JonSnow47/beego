package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/travel/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	//"github.com/JonSnow47/beego/application/travel/common"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	beego.Controller
}

func (a *AdminController) New() {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug("ErrInvalidParam", err)
		a.Data["json"] = map[string]interface{}{"status": "ErrInvalidParam"}
	} else {
		err := models.AdminServer.New(admin.Name, admin.Password)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": "No Data"}
		} else if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{"status": "Mysql error"}
		} else {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": admin.Name}
		}
	}
	a.ServeJSON()
}

func (a *AdminController) Login() {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug("ErrInvalidParam", err)
		a.Data["json"] = map[string]interface{}{"status": "ErrInvalidParam"}
	} else {
		err := models.AdminServer.Login(admin.Name, admin.Password)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": "No Admin"}
		} else if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{"status": "Mysql error"}
		} else {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": admin.Name}
		}
	}
	a.ServeJSON()
}

func (a *AdminController) Logout() {

}
