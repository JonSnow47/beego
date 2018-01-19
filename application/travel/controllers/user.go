package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/travel/common"
	"github.com/JonSnow47/beego/application/travel/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Register() {
	var user models.User
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Debug("Unmarshal", err)
		u.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.UserServer.Register(user.Name, user.Pass)
		if err != nil {
			logs.Debug(err)
			u.Data["json"] = map[string]interface{}{common.RespKeyStatus: err}
		} else {
			u.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed,
				"id":   user.ID,
				"name": user.Name}
		}
	}
	u.ServeJSON()
}

func (u *UserController) Login() {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Debug("Unmarshal", err)
		u.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		flag, err := models.UserServer.Login(user.Username, user.Password)
		if err != nil {
			if err == orm.ErrNoRows {
				logs.Debug(err)
				u.Data["json"] = map[string]interface{}{common.RespKeyStatus: err}
			} else {
				logs.Debug(err)
				u.Data["json"] = map[string]interface{}{common.RespKeyStatus: err}
			}
		} else {
			if flag {
				u.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			} else {
				u.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrPassword}
			}
		}
	}
}

func (this *UserController) UserInfo() {
	var Id struct {
		id int64
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
		this.Data["json"] = map[string]interface{}{"result": "delete success"}
	}
	this.ServeJSON()
}

func (a *UserController) Logout() {
	a.DestroySession()
	a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}

	a.ServeJSON()
}
