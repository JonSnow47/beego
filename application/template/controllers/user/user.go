package user

import (
	"github.com/astaxie/beego"
	"golang.org/x/tools/cmd/guru/testdata/src/alias"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Hello() {
	u.Data["json"] = map[string]string{"content": "Hello beego!"}
	u.ServeJSON()
}
