package tui

import "os"

type model struct {
	files    []os.DirEntry
	path     string
	position int
	width    int
	height   int
}

func InitialModel() model {
	return model{}
}
