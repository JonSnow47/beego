package routers

import (
	"github.com/astaxie/beego"
	"github.com/JonSnow47/beego/application/template/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/object", &controllers.ObjectController{})
}