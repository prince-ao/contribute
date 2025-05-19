package main

import (
	"fmt"
	"os"

	"github.com/prince-ao/contribute/internal/components"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/prince-ao/contribute/internal/utils"
)

func main() {
	p := tea.NewProgram(components.InitMainModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

type model struct {
	name textinput.Model
}

func initModel() model {
	ti := textinput.New()
	ti.Placeholder = "github_pat_xyz"
	ti.Focus()
	ti.Width = 20

	return model{
		name: ti,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			return m, func() tea.Msg { return utils.SwitchToYearComponent{} }
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	}

	m.name, cmd = m.name.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf("Enter your GitHub?\n\n%s\n\n%s", m.name.View(), "(esc to quit)") + "\n"
}
