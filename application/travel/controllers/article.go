package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/travel/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) New() {
	var article models.Article
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug("ErrInvalidParam", err)
		a.Data["json"] = map[string]interface{}{"status": "ErrInvalidParam"}
	} else {
		err := models.ArticleServer.New(article.Title, article.Content)
		if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{"status": "Mysql error"}
		} else {
			a.Data["json"] = map[string]interface{}{"status": "Success"}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Read() {
	var article struct{ Id int64 }
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug("ErrInvalidParam", err)
		a.Data["json"] = map[string]interface{}{"status": "ErrInvalidParam"}
	} else {
		article, err := models.ArticleServer.Read(article.Id)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": "No article"}
		} else if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{"status": "Mysql error"}
		} else {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": article}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Update() {

}

func (a *ArticleController) Delete() {

}

func (a *ArticleController) View() {
	var article struct{ Id int64 }
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug("ErrInvalidParam", err)
		a.Data["json"] = map[string]interface{}{"status": "ErrInvalidParam"}
	} else {
		article, err := models.ArticleServer.View(article.Id)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": "No article"}
		} else if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{"status": "Mysql error"}
		} else {
			a.Data["json"] = map[string]interface{}{"status": "Success", "data": article}
		}
	}
	a.ServeJSON()
}
