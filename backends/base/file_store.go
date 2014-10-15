package base

type FileStore struct{}

func (this *FileStore) ReadDirRecursive(path string, depth int) (paths []string, err error) {
	panic("Not implemented")
	return
}
