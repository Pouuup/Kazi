package tui

import (
	"os"

	"github.com/Pouuup/Kazi/internal/fs"
)

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
	m.files, m.err = fs.SearchDirectory(m.path)
	return m
}
