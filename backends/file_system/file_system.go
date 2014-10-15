package file_system

import (
	"io/ioutil"
	"os"

	"github.com/Marktown/frontend/backends/abstract"
)

type FileStore struct {
	abstract.FileStore
	FilePath string
}

func (this *FileStore) CreateFile(path string, content []byte) (err error) {
	err = ioutil.WriteFile(path, content, os.ModePerm)
	return
}

func (this *FileStore) ReadFile(path string) (content []byte, err error) {
	content, err = ioutil.ReadFile(path)
	return
}
