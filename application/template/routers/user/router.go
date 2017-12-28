package routers

import (
	_ "github.com/JonSnow47/beego/application/template/controllers"
	"github.com/JonSnow47/beego/application/template/controllers/user"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user", &user.UserController{})
	beego.Router("/user/login", &user.UserController{}, "post:Login")
	beego.Router("/user/logout", &user.UserController{}, "post:Logout")
	beego.Router("/user/register", &user.UserController{}, "post:Register")
	beego.Router("/user/userinfo", &user.UserController{}, "post:Info")

}
