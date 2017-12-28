package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)
type MainController struct {
    beego.Controller
}

func (this *MainController) Get() {
	u := this.Ctx.Request.FormValue("user")
	fmt.Println("hjkl:",u)
    this.Ctx.ResponseWriter.Write([]byte(`
         <html>
           <head>
             <title>Chat</title>
           </head>
           <body>
             Let's chat!
           </body>
</html>`))
    this.ServeJSON()
}
