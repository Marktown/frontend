package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Prepare () {
	this.Layout = "layouts/default.html.tpl"
}

func (this *HomeController) Get () {
	this.TplNames = "home/index.html.tpl"
}
