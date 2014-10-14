package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (this *BaseController) SetupEnv() {
	this.Data["RunMode"] = beego.RunMode
}
