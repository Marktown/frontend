package file_system

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"github.com/Marktown/frontend/backends/base"
)

type FileStore struct {
	base.FileStore
	FilePath string
}

func (this *FileStore) CreateFile(path string, reader io.Reader) (err error) {
	dataInBytes := []byte{}
	_, err = reader.Read(dataInBytes)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, dataInBytes, os.ModePerm)
	return
}

func (this *FileStore) UpdateFile(path string, reader io.Reader) (err error) {
	dataInBytes := []byte{}
	_, err = reader.Read(dataInBytes)
	if err != nil {
		return
	}
	err = ioutil.WriteFile(path, dataInBytes, os.ModePerm)
	return
}

func (this *FileStore) CreateDir(path string) (err error) {
	err = os.Mkdir(path, os.ModePerm)
	return
}

func (this *FileStore) ReadFile(path string) (reader io.Reader, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	reader = bufio.NewReader(file)
	return
}

func (this *FileStore) ReadDir(path string) (list []os.FileInfo, err error) {
	list, err = ioutil.ReadDir(path)
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
