package models

type FileStoreInterface interface {
  // create new file at path
  CreateFile(path string, content []byte) (written bool, err error)

  // update existing file at path
  UpdateFile(path string, content []byte) (written bool, err error)

  // create dir at path
  CreateDir(path string) (created bool, err error)

  // read content of file at path
  ReadFile(path string) (content []byte, err error)

  // list direct child paths within dir at path
  ReadDir(path string) (paths []string, err error)

  // move file or dir at path
  Move(path string, newPath string) (moved bool, err error)

  // delete file or dir at path
  Delete(path string) (deleted bool, err error)
}
