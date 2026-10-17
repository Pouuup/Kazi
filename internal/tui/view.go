package tui

import (
	"fmt"
	"strings"
)

func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	var b strings.Builder
	for i, val := range m.files {
		cursor := "  "
		if i == m.position {
			cursor = "> "
		}
		kind := "[is File]"
		if val.IsDir() {
			kind = "[is Dir]"
		}
		fmt.Fprintf(&b, "%s%s  %s\n\n", cursor, val.Name(), kind)
	}
	return b.String()
}

/*
func (m model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	totals := ""

	elements, err := fs.SearchDirectory(m.path)
	if err != nil {
		return fmt.Sprintln(err)
	}

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
*/
