package models

type FileHandlerInterface interface {
	Create(file *File)
	Update(file *File)
	Delete(file *File)
}
