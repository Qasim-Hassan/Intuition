package main

import (
	"fmt"
	"os"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	newMessageField textinput.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl+c", "ctrl+q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() tea.View {

	var style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).Padding(0, 2, 0, 2)

	welcomemsg := style.Render("Welcome to Intuition")
	help := "Ctrl+N: new file - Ctrl+L: list - Esc: back/save - Ctrl+S: save - Ctrl+Q: quit"

	view := ""

	return tea.View{Content: fmt.Sprintf("\n%s\n\n%s\n\n%s", welcomemsg, view, help)}
}

func initializeMode() model {
	return model{}
}

func main() {
	p := tea.NewProgram(initializeMode())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occured: %v", err)
		os.Exit(1)
	}
}
