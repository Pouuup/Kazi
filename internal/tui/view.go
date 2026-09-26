package tui

import "github.com/Pouuup/Kazi/internal/fs"

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	totals := ""
	elements, _ := fs.SearchDirectory(m.path)
	for i, val := range elements {
		if i == m.position {
			totals += "> "
		}

		if val.IsDir() {
			totals += val.Name() + "  [is Dir]"
			totals += "\n\n"
		} else {
			totals += val.Name() + "  [is File]"
			totals += "\n\n"
		}
	}

	return totals
}
