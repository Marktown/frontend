package models

import (
	"io/ioutil"
)

type FSFileStore struct {
}

func (this *FSFileStore) CreateFile(path string, content []byte) (created bool, err error) {
	// execute "touch " + file.Path
	// execute "echo \"" + file.Content "\" > " + file.Path
	return
}

func (this *FSFileStore) UpdateFile(path string, content []byte) (created bool, err error) {
	// execute "touch " + file.Path
	// execute "echo \"" + file.Content "\" > " + file.Path
	return
}

func (this *FSFileStore) CreateDir(path string) (created bool, err error) {
	// execute "touch " + file.Path
	// execute "echo \"" + file.Content "\" > " + file.Path
	return
}

func (this *FSFileStore) ReadFile(path string) (content []byte, err error) {
	content, err = ioutil.ReadFile(path)
	return
}

func (this *FSFileStore) ReadDir(path string) (paths []string, err error) {
	return
}

func (this *FSFileStore) Move(path string, newPath string) (moved bool, err error) {
	// execute "echo \"" + file.Content "\" > " + file.Path
	return
}

func (this *FSFileStore) Delete(path string) (deleted bool, err error) {
	// execute "rm " + file.Path
	return
}
