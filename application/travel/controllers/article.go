package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/Fan/common"
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
		logs.Debug("Unmarshal", err)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.New(article, a.GetSession(common.SessionUserID).(string))
		if err != nil {
			logs.Debug(err)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Read() {
	var article struct{ Id int64 }
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug("Unmarshal", err)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		article, err := models.ArticleServer.Read(article.Id)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: common.ErrNoData}
		} else if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: article}
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
		logs.Debug("Unmarshal", err)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		article, err := models.ArticleServer.View(article.Id)
		if err == orm.ErrNoRows {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: common.ErrNoData}
		} else if err != nil {
			logs.Debug("MysqlERR", err)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: article}
		}
	}
	a.ServeJSON()
}
