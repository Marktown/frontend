package models

type FSFileHandler struct {
}

func (this *FSFileHandler) Create(file *File) {
	// execute "touch " + file.Path
	// execute "echo \"" + file.Content "\" > " + file.Path
}

func (this *FSFileHandler) Update(file *File) {
	// execute "echo \"" + file.Content "\" > " + file.Path
}

func (this *FSFileHandler) Delete(file *File) {
	// execute "rm " + file.Path
}

func FileHandler() (fileHandler FileHandlerInterface) {
	return new(FSFileHandler)
}
