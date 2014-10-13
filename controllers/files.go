package controllers

import (
	"github.com/astaxie/beego"
)

type FilesController struct {
	beego.Controller
}

func (this *FilesController) Prepare () {
	this.Layout = "layouts/default.html.tpl"
}

func (this *FilesController) Index () {
	this.TplNames = "files/index.html.tpl"
}

func (this *FilesController) New () {
	this.TplNames = "files/new.html.tpl"
}

func (this *FilesController) Update () {
	this.TplNames = "files/update.html.tpl"
}

func (this *FilesController) Delete () {
	this.TplNames = "files/delete.html.tpl"
}
