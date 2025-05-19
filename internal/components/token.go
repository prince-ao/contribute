package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type tokenModel struct {
	name textinput.Model
}

func initTokenModel() tokenModel {
	ti := textinput.New()
	ti.Placeholder = "github_pat_xyz"
	ti.Focus()
	ti.Width = 20

	return tokenModel{
		name: ti,
	}
}

func (m tokenModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m tokenModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.name, cmd = m.name.Update(msg)
	return m, cmd
}

func (m tokenModel) View() string {
	return fmt.Sprintf("Enter your GitHub?\n\n%s\n\n%s", m.name.View(), "(esc to quit)") + "\n"
}
