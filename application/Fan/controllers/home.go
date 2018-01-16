package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"

	"github.com/JonSnow47/beego/application/Fan/common"
	"github.com/JonSnow47/beego/application/Fan/models"
)

type MainController struct {
	beego.Controller
}

func (h *MainController) Home() {
	var article struct {
		Title string
	}
	err := json.Unmarshal(h.Ctx.Input.RequestBody, &article)
	if err != nil {
		logs.Debug(err)
		h.Data["json"] = map[string]interface{}{common.RespKeyErr: err}
	} else {
		article, err := models.ArticleServer.View(article.Title)
		if err != nil {
			logs.Debug(err)
			h.Data["json"] = map[string]interface{}{common.RespKeyErr: common.ErrNoData}
		} else {
			h.Data["json"] = map[string]interface{}{common.RespKeyData: article}
		}
	}
	h.ServeJSON()
}
