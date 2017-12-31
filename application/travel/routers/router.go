package routers

import (
	"github.com/JonSnow47/beego/application/travel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user",&controllers.UserController{},"get:Hello")
	beego.Router("/user/id", &controllers.UserController{},"get:UserID")

}
