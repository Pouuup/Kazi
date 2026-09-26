package fs

import (
	"os"
)

func SearchDirectory(currentPath string) ([]os.DirEntry, error) {
	d, err := os.ReadDir(currentPath)
	if err != nil {
		return nil, err
	}
	total := []os.DirEntry{}
	for _, val := range d {
		total = append(total, val)
	}
	return total, nil

}
