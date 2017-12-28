package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct{
	beego.Controller
}

func (this *MainController) Get()  {
	this.Ctx.ResponseWriter.Write([]byte(`
         <html>
           <head>
             <title>Chat</title>
           </head>
           <body>
             Welcome to beego!
           </body>
	</html>`))
	this.ServeJSON()
}