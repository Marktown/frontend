package models

type FileHandlerInterface interface {
	Create(file *File)
	Read(file *File) (string, error)
	Update(file *File)
	Delete(file *File)
}
