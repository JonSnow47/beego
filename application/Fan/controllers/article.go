package controllers

import (
	"encoding/json"
	"github.com/JonSnow47/beego/application/Fan/models"

	"github.com/JonSnow47/beego/application/Fan/common"
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
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.New(article.Title, article.Author, article.Class, article.Content)
		if err != nil {
			a.Data["json"] = "can not create article"
		} else {
			//info,_ := models.AdminServer.AdminInfo(admin.Name)
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: article.Title}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Read() {
	var Id struct {
		Id int64
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &Id)
	if err != nil {
		logs.Debug(err)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		article, err := models.ArticleServer.Read(Id.Id)
		if err != nil {
			logs.Debug("ArticleServer.Read:", err)
		} else {
			a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.RespKeyData: article}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Update() {
	massage := models.Display{}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &massage)
	if err != nil {
		logs.Debug(massage)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Update(massage)
		if err != nil {
			a.Data["json"] = "Update error."
		} else {
			a.Data["json"] = map[string]interface{}{"Update success": massage}
		}
	}
	a.ServeJSON()
}

func (a *ArticleController) Delete() {
	var massage struct {
		Title  string
		Author string
	}
	err := json.Unmarshal(a.Ctx.Input.RequestBody, &massage)
	if err != nil {
		logs.Debug(massage)
		a.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.ArticleServer.Delete(massage.Title, massage.Author)
		if err != nil {
			a.Data["json"] = "Delete error."
		} else {
			a.Data["json"] = "Delete success."
		}
	}
	a.ServeJSON()
}
