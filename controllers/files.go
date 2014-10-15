package controllers

import (
	"github.com/Marktown/frontend/backends/file_system"
	"text/scanner"
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
	fs := file_system.NewFileStore()
	fs.RootPath = "tests/assets/testfolder/"
	list, err := fs.ReadDir("")
	//TODO error handling
	if err != nil {
		return
	}
	//TODO use ReadDirTree for recursive reading, must be implemented
	pathItems := PathItem{"/", []PathItem{}}
	items := make([]PathItem, len(list))
	for index, file := range list {
		items[index] = PathItem{file.Name(), []PathItem{}}
	}
	pathItems.Items = items
	this.Data["json"] = &pathItems
	this.ServeJson()
}

func (this *FilesController) New() {
	this.TplNames = "files/new.html.tpl"
}

func (this *FilesController) Read() {
	textfile := Textfile{}
	fs := file_system.NewFileStore()
	fs.RootPath = "tests/assets/testfolder/"
	reader, err := fs.ReadFile(this.GetString("path"))
	if err != nil {
		return
	}

	textfile.Data = ""
	var s scanner.Scanner
	s.Init(reader)
	s.Whitespace = 1
	tok := s.Scan()
	for tok != scanner.EOF {
		textfile.Data += s.TokenText()
		tok = s.Scan()
	}
	this.Data["json"] = &textfile
	this.ServeJson()
}

func (this *FilesController) Update() {
	this.TplNames = "files/update.html.tpl"
}

func (this *FilesController) Delete() {
	this.TplNames = "files/delete.html.tpl"
}
