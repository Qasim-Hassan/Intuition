package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	msg string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// Is it a key press?
	case tea.KeyPressMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() tea.View {
	return tea.View{Content: m.msg}
}

func initializeMode() model {
	return model{
		msg: "ayo",
	}
}

func main() {
	p := tea.NewProgram(initializeMode())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occured: %v", err)
		os.Exit(1)
	}
}
