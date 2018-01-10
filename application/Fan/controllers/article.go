package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/Fan/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) New() {
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
			a.Data["json"] = map[string]interface{}{"Successful": article}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Read() {
	var id int64
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &id)
	if err != nil {
		logs.Debug(err)
		a.Data["json"] = "Request error."
	} else {
		article, err := models.ArticleServer.Read(id)
		if err != nil {
			logs.Debug(err)
			a.Data["json"] = map[string]interface{}{"Successful": article}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Updata() {
	/*type massage struct {
		Title	string
		Class	string
		Content	string
	}*/
	massage := models.Display{}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &massage)
	if err != nil {
		logs.Debug(massage)
		a.Data["json"] = "Request error."
	} else {
		err := models.ArticleServer.Updata(massage)
		if err != nil {
			a.Data["json"] = "Updata error."
		} else {
			a.Data["json"] = map[string]interface{}{"Updata success": massage}
		}
	}
	a.ServeJSON()
}
func (a *ArticleController) Delete() {
	type massage struct {
		Title  string
		Auther string
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &massage{})
	if err != nil {
		logs.Debug(massage{})
		a.Data["json"] = "Request error."
	} else {
		err := models.ArticleServer.Delete(massage{}.Title, massage{}.Auther)
		if err != nil {
			a.Data["json"] = "Delete error."
		} else {
			a.Data["json"] = "Delete success."
		}
	}
	a.ServeJSON()
}
