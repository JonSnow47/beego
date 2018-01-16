package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"

	"github.com/JonSnow47/beego/application/Fan/common"
	"github.com/JonSnow47/beego/application/Fan/models"
)

var globalSessions *session.Manager

type AdminController struct {
	beego.Controller
}

//insert a new admin
func (a *AdminController) New() {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug(common.ErrInvalidParam, admin)
		a.Data["json"] = map[string]interface{}{common.RespKeyErr: err}
	} else {
		err := models.AdminServer.Create(admin.Name, admin.Password)
		if err != nil {
			logs.Debug(common.ErrMysqlQuery, err)
			a.Data["json"] = map[string]interface{}{common.RespKeyErr: err}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyData: admin.Name}
		}
	}
	a.ServeJSON()
}

//check admin info
func (a *AdminController) Login() {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug(common.ErrInvalidParam, admin)
		a.Data["json"] = map[string]interface{}{common.RespKeyErr: common.ErrInvalidParam}
	} else {
		ok, err := models.AdminServer.Login(admin.Name, admin.Password)
		if err != nil {
			logs.Debug(common.ErrMysqlQuery, err)
			a.Data["json"] = map[string]interface{}{common.RespKeyErr: common.ErrMysqlQuery}
		} else if err == nil && ok == true {
			a.SetSession(common.SessionUserID, admin.Name)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrLogin}
		}
	}
	a.ServeJSON()
}

func (a *AdminController) Logout() {
	a.DestroySession()
	a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}

	a.ServeJSON()
}
