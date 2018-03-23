package routers

import (
	"github.com/JonSnow47/beego/application/pampas/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
