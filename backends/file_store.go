package backends

type FileStore interface {
	// create new file at path
	CreateFile(path string, content []byte) (err error)

	// update existing file at path
	UpdateFile(path string, content []byte) (err error)

	// create dir at path
	CreateDir(path string) (err error)

	// read content of file at path
	ReadFile(path string) (content []byte, err error)

	// list direct child paths within dir at path
	ReadDir(path string) (paths []string, err error)

	// list direct and indirect child paths within dir at path for a given depth
	// depth -1 means unlimited depth
	ReadDirTree(path string, depth int) (paths []string, err error)

	// move file or dir at path
	Move(path string, newPath string) (err error)

	// delete file or dir at path
	Delete(path string) (err error)
}
