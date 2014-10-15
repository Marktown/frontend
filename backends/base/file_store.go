package base

import (
	"os"
)

type FileStore struct{}

func (this *FileStore) ReadDirTree(path string, depth int) (list []os.FileInfo, err error) {
	panic("not implemented")
	return
}
