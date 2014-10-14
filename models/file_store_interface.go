package models

type FileStoreInterface interface {
  CreateFile(path string, content []byte) (written bool, err error)
  UpdateFile(path string, content []byte) (written bool, err error)
  CreateDir(path string) (created bool, err error)
  ReadFile(path string) (content []byte, err error)
  ReadDir(path string) (paths []string, err error)
  Move(path string, newPath string) (moved bool, err error)
  Delete(path string) (deleted bool, err error)
}
