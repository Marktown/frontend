package backends

import (
	"os"
)

type FileStore interface {
	// create new file at given path
	CreateFile(path string, content []byte) (err error)

	// update existing file at given path
	UpdateFile(path string, content []byte) (err error)

	// create dir at given path
	CreateDir(path string) (err error)

	// read content of file at given path
	ReadFile(path string) (content []byte, err error)

	// list direct child paths within dir at path
	ReadDir(path string) (list []os.FileInfo, err error)

	// list direct and indirect childs within dir at given path for a given depth
	// depth -1 means unlimited depth
	ReadDirTree(path string, depth int) (paths []File, err error)

	// move file or dir at given path
	Move(path string, newPath string) (err error)

	// delete file or dir at given path
	Delete(path string) (err error)
}
