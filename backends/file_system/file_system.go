package file_system

import (
  "io/ioutil"
)

type FileStore struct {
}

func (this *FileStore) CreateFile(path string, content []byte) (created bool, err error) {
  // execute "touch " + file.Path
  // execute "echo \"" + file.Content "\" > " + file.Path
  return
}

func (this *FileStore) UpdateFile(path string, content []byte) (created bool, err error) {
  // execute "touch " + file.Path
  // execute "echo \"" + file.Content "\" > " + file.Path
  return
}

func (this *FileStore) CreateDir(path string) (created bool, err error) {
  // execute "touch " + file.Path
  // execute "echo \"" + file.Content "\" > " + file.Path
  return
}

func (this *FileStore) ReadFile(path string) (content []byte, err error) {
  content, err = ioutil.ReadFile(path)
  return
}

func (this *FileStore) ReadDir(path string) (paths []string, err error) {
  return
}

func (this *FileStore) Move(path string, newPath string) (moved bool, err error) {
  // execute "echo \"" + file.Content "\" > " + file.Path
  return
}

func (this *FileStore) Delete(path string) (deleted bool, err error) {
  // execute "rm " + file.Path
  return
}
