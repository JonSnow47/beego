package controllers

import (
	"github.com/astaxie/beego"

	"github.com/JonSnow47/beego/application/Fan/models"

	"encoding/json"
	"github.com/astaxie/beego/logs"
)

type MainController struct {
	beego.Controller
}

func (h *MainController) Home()  {
	var id int64
	err := json.Unmarshal(h.Ctx.Input.RequestBody,&id)
	if err != nil {
		logs.Debug(err)
		h.Data["json"] = "Request error."
	} else {
		article,err := models.ArticleServer.Read(id)
		if err != nil {
			logs.Debug(err)
			h.Data["json"] = map[string]interface{}{"Successful":article}
			}
		h.ServeJSON()
		}
}