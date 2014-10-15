package controllers

import (
	"io/ioutil"
)

type PathItem struct {
	Path  string     `json:"name"`
	Items []PathItem `json:"items"`
}

type Textfile struct {
	Data string `json:"data"`
}

type FilesController struct {
	BaseController
}

func (this *FilesController) Prepare() {
	this.SetupEnv()
	this.Layout = "layouts/default.html.tpl"
}

func (this *FilesController) Index() {
	this.TplNames = "files/index.html.tpl"
	pathItems := PathItem{"/", []PathItem{
		PathItem{"/foo", []PathItem{}},
		PathItem{"/bar", []PathItem{
			PathItem{"/bar/bax", []PathItem{}},
			PathItem{"/bar/baz", []PathItem{}},
		}},
		PathItem{"/fax", []PathItem{
			PathItem{"/fax/ban", []PathItem{}},
			PathItem{"/fax/bam", []PathItem{}},
		}},
		PathItem{"/faz", []PathItem{}},
	}}
	this.Data["json"] = &pathItems
	this.ServeJson()
}

func (this *FilesController) New() {
	this.TplNames = "files/new.html.tpl"
}

func (this *FilesController) Read() {
	textfile := Textfile{}
	dataAsBytes, err := ioutil.ReadFile("tests/assets/testfile.md")
	if err != nil {
		return
	}
	textfile.Data = string(dataAsBytes)
	this.Data["json"] = &textfile
	this.ServeJson()
}

func (this *FilesController) Update() {
	this.TplNames = "files/update.html.tpl"
}

func (this *FilesController) Delete() {
	this.TplNames = "files/delete.html.tpl"
}
