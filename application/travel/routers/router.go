package routers

import (
	"github.com/JonSnow47/beego/application/travel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	//user
	beego.Router("/user", &controllers.UserController{}, "*:Hello")
	beego.Router("/user/register", &controllers.UserController{}, "post:Register")
	beego.Router("/user/info", &controllers.UserController{}, "get:UserInfo")
}
