package controllers

import (
	"github.com/Marktown/frontend/models"
)

type FilesController struct {
	BaseController
}

func (this *FilesController) Prepare() {
	this.SetupEnv()
	this.Layout = "layouts/default.html.tpl"
}

func (this *FilesController) Index() {
	this.TplNames = "files/index.html.tpl"
}

func (this *FilesController) New() {
	this.TplNames = "files/new.html.tpl"
	file := new(models.File)
	fileHandler := new(models.FileHandler)
	fileHandler.Create(file)
}

func (this *FilesController) Update() {
	this.TplNames = "files/update.html.tpl"
	file := new(models.File)
	fileHandler := new(models.FileHandler)
	fileHandler.Update(file)
}

func (this *FilesController) Delete() {
	this.TplNames = "files/delete.html.tpl"
	file := new(models.File)
	fileHandler := new(models.FileHandler)
	fileHandler.Remove(file)
}
