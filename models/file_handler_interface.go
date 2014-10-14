package models

type FileHandlerInterface interface {
	Create(file *File) error
	Read(file *File) (string, error)
	Update(file *File) error
	Delete(file *File) error
}
