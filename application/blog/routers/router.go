
package routers

import (
	"github.com/JonSnow47/beego/application/blog/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/admin/new", &controllers.AdminController{}, "Post:New")

	beego.Router("/article/new", &controllers.ArticleController{}, "Post:New")
}
