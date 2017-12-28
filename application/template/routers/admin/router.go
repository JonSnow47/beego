package routers

import (
	"github.com/astaxie/beego"
	_ "github.com/JonSnow47/beego/application/template/controllers"
	"github.com/JonSnow47/beego/application/template/controllers/admin"
	)

func init()  {
	beego.Router("/admin",&admin.AdminController{},"post:Change")

}