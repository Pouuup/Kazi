package fs

import (
	"fmt"
	"os"
)

func SearchDirectory(currentPath string) []os.DirEntry {
	d, err := os.ReadDir(currentPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	total := []os.DirEntry{}
	for _, val := range d {
		total = append(total, val)
	}
	return total

}
