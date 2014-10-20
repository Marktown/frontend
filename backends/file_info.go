package backends

type FileInfo interface {
	Path string
	IsDir bool
	// size?
	// permissions?
}
