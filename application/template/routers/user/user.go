package routers

import (
	"github.com/astaxie/beego"
	"github.com/JonSnow47/beego/application/template/controllers"
)

func init()  {
	beego.Router("/user",&controllers.UserController{})
	beego.Router("/user/login",&controllers.UserController{},"post:Login")
	beego.Router("/user/logout",&controllers.UserController{},"post:Logout")

	}