package controllers

import (
	"github.com/JonSnow47/beego/application/Fan/models"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) New()  {
	var article models.Article
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug(article)
		a.Data["json"] = "Request error"
	} else {
		err := models.ArticleServer.New(article)
		if err != nil {
			a.Data["json"] = "can not create article"
		} else {
			//info,_ := models.AdminServer.AdminInfo(admin.Name)
			a.Data["json"] = map[string]string{"info":article}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Read()  {

}

func (a *ArticleController) Updata () {
	var article models.Article
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug(article)
		a.Data["json"] = "Request error."
	} else {

	}
}
func (a *ArticleController) Delete()  {

}