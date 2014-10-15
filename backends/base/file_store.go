package base

import (
	"github.com/Marktown/frontend/backends"
)

type FileStore struct{}

func (this *FileStore) ReadDirTree(path string, depth int) (paths []backends.File, err error) {
	panic("Not implemented")
	return
}
