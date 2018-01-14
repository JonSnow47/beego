package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/JonSnow47/beego/application/Fan/common"
	"github.com/JonSnow47/beego/application/Fan/models"
	"github.com/astaxie/beego/orm"
)

type ArticleController struct {
	beego.Controller
}

func (a *ArticleController) New() {
	var article models.Article
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug(common.ErrInvalidParam, article)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.New(article.Title, article.Class, article.Content)
		if err != nil {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			id := models.ArticleServer.GetArticleID(article.Title)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: id}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Read() {
	var article struct {
		Id int64
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug(common.ErrInvalidParam, err)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		article, err := models.ArticleServer.Read(article.Id)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrNoData}
		} else if err != nil {
			logs.Debug(common.ErrMysqlQuery, err)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: article}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Update() {
	var article models.Article
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug(common.ErrInvalidParam, article)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Update(article)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrNoData}
		} else if err != nil {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Delete() {
	var article struct {
		Title string
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug(common.ErrInvalidParam, article)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		id := models.ArticleServer.GetArticleID(article.Title)
		err := models.ArticleServer.Delete(id)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrNoData}
		} else if err != nil {
			logs.Debug(common.ErrMysqlQuery, err)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	a.ServeJSON()
}
