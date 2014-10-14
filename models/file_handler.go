package models

type FileHandler struct {
}

func (this *FileHandler) Create (file *File) {
	// execute "touch " + file.Path
	// execute "echo \"" + file.Content "\" > " + file.Path
}

func (this *FileHandler) Update (file *File) {
	// execute "echo \"" + file.Content "\" > " + file.Path
}

func (this *FileHandler) Remove (file *File) {
	// execute "rm " + file.Path
}
