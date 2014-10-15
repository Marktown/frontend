package file_system

import (
	"io/ioutil"
	"os"

	"github.com/Marktown/frontend/backends"
	"github.com/Marktown/frontend/backends/base"
)

type FileStore struct {
	base.FileStore
	FilePath string
}

func (this *FileStore) CreateFile(path string, content []byte) (err error) {
	err = ioutil.WriteFile(path, content, os.ModePerm)
	return
}

func (this *FileStore) UpdateFile(path string, content []byte) (err error) {
	panic("Not implemented")
	return
}

func (this *FileStore) CreateDir(path string) (err error) {
	panic("Not implemented")
	return
}

func (this *FileStore) ReadFile(path string) (content []byte, err error) {
	content, err = ioutil.ReadFile(path)
	return
}

func (this *FileStore) ReadDir(path string) (paths []backends.File, err error) {
	panic("Not implemented")
	return
}

func (this *FileStore) Move(path string, newPath string) (err error) {
	panic("Not implemented")
	return
}

func (this *FileStore) Delete(path string) (err error) {
	panic("Not implemented")
	return
}
