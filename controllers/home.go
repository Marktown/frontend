package controllers

type HomeController struct {
	BaseController
}

func (this *HomeController) Prepare() {
	this.SetupEnv()
	this.Layout = "layouts/default.html.tpl"
}

func (this *HomeController) Get() {
	this.TplNames = "home/index.html.tpl"
}
