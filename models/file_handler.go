package models

import (
	"io/ioutil"
)

type FSFileHandler struct {
}

func (this *FSFileHandler) Create(file *File) {
	// execute "touch " + file.Path
	// execute "echo \"" + file.Content "\" > " + file.Path
}

func (this *FSFileHandler) Read(file *File) (content string, err error) {
	buffer, err := ioutil.ReadFile(file.Path)
	if err == nil {
		content = string(buffer)
	}
	return
}

func (this *FSFileHandler) Update(file *File) {
	// execute "echo \"" + file.Content "\" > " + file.Path
}

func (this *FSFileHandler) Delete(file *File) {
	// execute "rm " + file.Path
}

func FileHandler() (fileHandler FileHandlerInterface) {
	return new(FSFileHandler)
}
