package tui

import "os"

type model struct {
	files    []os.DirEntry
	path     string
	err      error
	position int
	width    int
	height   int
}

func InitialModel() model {
	m := model{}
	var err error
	m.path, err = os.Getwd()
	if err != nil {
		return model{err: err}
	}

	return m
}
