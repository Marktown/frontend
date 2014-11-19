package backends

type FileInfo interface {
	Path() string
	IsDir() bool
	Name() string // return only file name
	// size?
	// permissions?
}
