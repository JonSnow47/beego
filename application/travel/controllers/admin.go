package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/JonSnow47/beego/application/travel/common"
	"github.com/JonSnow47/beego/application/travel/models"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	beego.Controller
}

func (a *AdminController) New() {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug("Unmarshal", err)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.AdminServer.New(admin.Name, admin.Password)
		if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrAdminExists}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: admin.Name}
		}
	}
	a.ServeJSON()
}

func (a *AdminController) Login() {
	var admin struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug("Unmarshal", err)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		flag, err := models.AdminServer.Login(admin.Name, admin.Password)
		if err != nil {
			if err == orm.ErrNoRows {
				a.Data["json"] = map[string]string{common.RespKeyStatus: common.ErrInvalidAdmin}
			} else {
				a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery, "error": err}
			}
		} else {
			if flag {
				a.Data["json"] = map[string]string{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: admin.Name}
			} else {
				a.Data["json"] = map[string]string{common.RespKeyStatus: common.ErrPassword}
			}
		}
	}
	a.ServeJSON()
}

func (a *AdminController) Logout() {
	a.DestroySession()
	a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}

	a.ServeJSON()
}
