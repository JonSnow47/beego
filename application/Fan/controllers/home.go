package controllers

import (
	"github.com/astaxie/beego"

	"github.com/JonSnow47/beego/application/Fan/models"

	"encoding/json"
)

type MainController struct {
	beego.Controller
}

func (h *MainController) Home()  {
	var article models.Article
	article = models.Hot
	h.Data["json"] = map[string]interface{}{article}
}