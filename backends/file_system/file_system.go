package file_system

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"

	"bytes"
	"crypto/sha1"

	"github.com/Marktown/frontend/backends"
)

type FileStore struct {
	RootPath string
}

func NewFileStore() *FileStore {
	fs := &FileStore{}
	fs.RootPath = "tmp/fs_file_store/"
	return fs
}

func (this *FileStore) WriteFile(path string, reader io.Reader) (err error) {
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

func (this *FileStore) ReadDir(path string) (list []backends.FileInfo, err error) {
	infoList, err := ioutil.ReadDir(this.RootPath + path)
	if err != nil {
		return
	}
	list = []backends.FileInfo{}
	for index, info := range infoList {
		list[index] = NewFileInfo(this.RootPath, info)
	}
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

func (this *FileStore) Checksum(path string) (sum [sha1.Size]byte, err error) {
	reader, err := this.ReadFile(path)
	if err != nil {
		return
	}
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	sum = sha1.Sum(content)
	return
}

func (this *FileStore) ReadDirTree(path string, depth int) (list []backends.FileInfo, err error) {
	list, err = this.ReadDirTreeHelper(path, depth, []backends.FileInfo{})
	return
}

func (this *FileStore) ReadDirTreeHelper(path string, depth int, tmpList []backends.FileInfo) (list []backends.FileInfo, err error) {
	infoList, err := this.ReadDir(path)
	if err != nil {
		return
	}
	for _, info := range infoList {
		tmpList = append(tmpList, []backends.FileInfo{info}...)
		if info.IsDir() {
			depthCopy := depth - 1
			if depthCopy > 0 {
				tmpList, err = this.ReadDirTreeHelper(path+"/"+info.Name(), depthCopy, tmpList)
			}
		}
	}
	list = tmpList
	return
}

func (this *FileStore) ReadRoot() (list []backends.FileInfo, err error) {
	return this.ReadDir("")
}
