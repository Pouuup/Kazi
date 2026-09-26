package tui

import (
	"path/filepath"

	"github.com/Pouuup/Kazi/internal/fs"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	elements, err := fs.SearchDirectory(m.path)
	if err != nil {
		m.err = err
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "k":
			if m.position > 0 {
				m.position--
			}
		case "down", "j":
			if m.position < len(elements)-1 {
				m.position++
			}
		case "enter":
			if len(elements) == 0 {
				return m, nil
			}

			current := elements[m.position]

			if !current.IsDir() {
				name := current.Name()
				pathCurrentFile := filepath.Join(m.path, name)

				return m.editingFile(pathCurrentFile)
			}

			m.path = filepath.Join(m.path, current.Name())

			m.position = 0
		case "esc":
			m.path = filepath.Dir(m.path)
			m.position = 0
		}
	}

	return m, nil
}
