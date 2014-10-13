package routers

import (
	"github.com/Marktown/frontend/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.HomeController{})

	beego.Router("/files", &controllers.FilesController{}, "get:Index")
	beego.Router("/files/new", &controllers.FilesController{}, "get:New")
	beego.Router("/files/update", &controllers.FilesController{}, "get:Update")
	beego.Router("/files/delete", &controllers.FilesController{}, "get:Delete")
}
