package tui

import (
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
)

type editorFinishMsg struct{ err error }

func (m model) editingFile(currentFile string) (model, tea.Cmd) {
	editing := exec.Command("nvim", currentFile)

	return m, tea.ExecProcess(editing, func(err error) tea.Msg {
		return editorFinishMsg{err}
	})
}
