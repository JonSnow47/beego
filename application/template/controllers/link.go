package controllers

import "github.com/astaxie/beego"

type LinkController struct {
	beego.Controller
}

func (this *LinkController) Get() {
	this.Ctx.Request.FormValue("aab")
	this.Ctx.ResponseWriter.Write([]byte(`
    <html>
           <head>
             <title>Links</title>
           </head>
           <body>
			http://github.com/JonSnow47
           </body>
	</html>`))
	this.ServeJSON()
}