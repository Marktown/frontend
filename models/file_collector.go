package models

type FSFileCollector struct {
}

func (this *FSFileCollector) FetchFiles(path string) (files []File) {
	// TODO fetch all files in the given path
	files = make([]File, 10)
	return files
}
