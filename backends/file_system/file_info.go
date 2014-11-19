package file_system

import (
	"os"
	"path"
)

type FileInfo struct {
	path  string
	isDir bool
}

func NewFileInfo(rootPath string, info os.FileInfo) *FileInfo {
	newInfo := &FileInfo{}
	newInfo.path = rootPath + info.Name()
	newInfo.isDir = info.IsDir()
	return newInfo
}

func (this *FileInfo) IsDir() bool {
	return this.isDir
}

func (this *FileInfo) Path() string {
	return this.path
}

func (this *FileInfo) Name() string {
	return path.Base(this.Path())
}
