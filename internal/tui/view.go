package tui

import "github.com/Pouuup/Kazi/internal/fs"

func (m model) View() string {
	totals := ""
	elements := fs.SearchDirectory(m.path)
	for i, val := range elements {
		if i == m.position {
			totals += "> "
		}
		totals += val.Name()
		totals += "\n\n"
	}

	return totals
}
