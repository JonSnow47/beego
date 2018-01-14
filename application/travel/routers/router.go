package routers

import (
	"github.com/JonSnow47/beego/application/travel/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.MainController{})

	//user
	beego.Router("/user", &controllers.UserController{}, "*:Hello")
	beego.Router("/user/register", &controllers.UserController{}, "post:Register")
	//beego.Router("/user/login", &controllers.UserController{}, "post:Login")
	//beego.Router("/user/logout", &controllers.UserController{}, "post:Logout")
	//admin
	beego.Router("/admin/new", &controllers.AdminController{}, "post:New")
	beego.Router("/admin/login", &controllers.AdminController{}, "post:Login")
	beego.Router("/admin/logout", &controllers.AdminController{}, "post:Logout")
	//article
	beego.Router("/article/new", &controllers.ArticleController{}, "post:New")
	beego.Router("/article/read", &controllers.ArticleController{}, "post:Read")
	beego.Router("/article/update", &controllers.ArticleController{}, "post:Update")
	beego.Router("/article/delete", &controllers.ArticleController{}, "post:Delete")
	beego.Router("/article/view", &controllers.ArticleController{}, "post:View")
}
