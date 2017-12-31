package controllers

import "github.com/astaxie/beego"

type MainController struct {
	beego.Controller
}

func(this *MainController) Get()  {
	this.Data["json"] = "Welcome to Travel!"
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