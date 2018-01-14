package routers

import (
	"github.com/astaxie/beego"

	"github.com/JonSnow47/beego/application/Fan/controllers"
)

func init() {
	//homepage
	beego.Router("/", &controllers.MainController{}, "*:Home")
	//admin
	beego.Router("/admin/new", &controllers.AdminController{}, "post:New")
	beego.Router("/admin/login", &controllers.AdminController{}, "post:Login")
	//beego.Router("/admin/logout", &controllers.AdminController{}, "post:Signout")

	//article
	beego.Router("/article/new", &controllers.ArticleController{}, "post:New")
	beego.Router("/article/read", &controllers.ArticleController{}, "post:Read")
	beego.Router("/article/update", &controllers.ArticleController{}, "post:Update")
	beego.Router("/article/delete", &controllers.ArticleController{}, "post:Delete")
	//beego.Router("/article/getid", &controllers.ArticleController{},"post:GetID")
}
