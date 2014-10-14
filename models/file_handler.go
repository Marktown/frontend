package models

import (
	"io/ioutil"
	"os"
)

type FSFileHandler struct {
	FileFolder string
}

func (this *FSFileHandler) Create(file *File) (err error) {
	err = ioutil.WriteFile(this.FileFolder + file.Path, []byte(file.Content), os.ModePerm)
	return
}

func (this *FSFileHandler) Read(file *File) (content string, err error) {
	buffer, err := ioutil.ReadFile(this.FileFolder + file.Path)
	if err == nil {
		content = string(buffer)
	}
	return
}

func (this *FSFileHandler) Update(file *File) (err error) {
	err = ioutil.WriteFile(this.FileFolder + file.Path, []byte(file.Content), os.ModePerm)
	return
}

func (this *FSFileHandler) Delete(file *File) (err error) {
	// TODO delete
	return
}

func FileHandler() (fileHandler FileHandlerInterface) {
	fh := new(FSFileHandler)
	fh.FileFolder = "assets/"
	return fh
}
