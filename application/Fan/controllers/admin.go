package controllers

import (
	"github.com/JonSnow47/beego/application/Fan/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// Operations about Users
type AdminController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (a *AdminController) New()  {
	var admin models.Admin
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &admin)
	if err != nil {
		logs.Debug(admin)
		a.Data["json"] = "Reques error"
	} else {
		err := models.AdminServer.New(admin)
		if err != nil {
			a.Data["json"] = "New false"
		} else {
			a.Data["json"] = map[string]string{"result":admin}
			}
	}

}

func (a *AdminController) Login() {
	var Logininfo struct{
		name	string
		password	string
		}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &Logininfo)
	if err != nil {
		logs.Debug(Logininfo)
		a.Data["json"] = logs.Debug
	} else {
		err := models.AdminServer.Login(Logininfo)
		if err := nil {
			logs.Debug("SQL error")
			a.Data["json"] = "SQL error"
		} else {

		}
	}
	a.ServeJSON()
}

func (a *AdminController) Logout()  {

}