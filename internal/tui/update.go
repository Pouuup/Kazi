package tui

import (
	"github.com/Pouuup/Kazi/internal/fs"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	m.path = "."
	elements := fs.SearchDirectory(m.path)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up", "j":
			if m.position > 0 {
				m.position--
			}
		case "down", "k":
			if m.position < len(elements)-1 {
				m.position++
			}
		case "enter":
			return m.editingFile("go.mod")
		}
	}

	return m, nil
}
