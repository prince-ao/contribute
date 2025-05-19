package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/prince-ao/contribute/internal/utils"
)

type mainModel struct {
	activeModel tea.Model
}

func InitMainModel() mainModel {
	return mainModel{
		activeModel: initTokenModel(),
	}
}

func (m mainModel) Init() tea.Cmd {
	return nil
}

func (m mainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.(type) {
	case utils.SwitchToTokenComponent:
		m.activeModel = initTokenModel()
		return m, m.activeModel.Init()
	}

	m.activeModel, cmd = m.activeModel.Update(msg)
	return m, cmd
}

func (m mainModel) View() string {
	return m.activeModel.View()
}
