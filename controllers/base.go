package controllers

import (
	"github.com/Marktown/frontend/utils"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SetupEnv() {
	this.Data["RunMode"] = beego.RunMode
	this.Data["context"] = utils.NewContext(this.GetControllerAndAction())
}
