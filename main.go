package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	msg string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m model) View() tea.View {
	return tea.View{Content: "baigan"}
}

func initializeMode() model {
	return model{
		msg: "ayo",
	}
}

func main() {
	fmt.Println("Welcome to Intuition")
}
