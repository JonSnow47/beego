package routers

import (
	"github.com/JonSnow47/beego/application/template/controllers"
	"github.com/JonSnow47/beego/application/template/controllers/user"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/user", &user.UserController{} , "get:Hello")
	beego.Router("/user/hello", &user.UserController{}, "get:Hello")
}
