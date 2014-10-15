package file_system

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"bytes"

	"github.com/Marktown/frontend/backends/base"
)

type FileStore struct {
	base.FileStore
	RootPath string
}

func NewFileStore() *FileStore {
	fs := &FileStore{}
	fs.RootPath = "tmp/fs_file_store/"
	return fs
}

func (this *FileStore) CreateFile(path string, reader io.Reader) (err error) {
	buffer := bytes.NewBuffer([]byte{})
	_, err = buffer.ReadFrom(reader)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(this.RootPath+path, buffer.Bytes(), os.ModePerm)
	return
}

func (this *FileStore) UpdateFile(path string, reader io.Reader) (err error) {
	buffer := bytes.NewBuffer([]byte{})
	_, err = buffer.ReadFrom(reader)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(this.RootPath+path, buffer.Bytes(), os.ModePerm)
	return
}

func (this *FileStore) CreateDir(path string) (err error) {
	err = os.Mkdir(this.RootPath+path, os.ModePerm)
	return
}

func (this *FileStore) ReadFile(path string) (reader io.Reader, err error) {
	file, err := os.Open(this.RootPath + path)
	if err != nil {
		return
	}
	reader = bufio.NewReader(file)
	return
}

func (this *FileStore) ReadDir(path string) (list []os.FileInfo, err error) {
	list, err = ioutil.ReadDir(this.RootPath + path)
	return
}

func (this *FileStore) Move(path string, newPath string) (err error) {
	err = os.Rename(this.RootPath+path, this.RootPath+newPath)
	return
}

func (this *FileStore) Delete(path string) (err error) {
	err = os.Remove(this.RootPath + path)
	return
}
