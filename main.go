package main

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

type model struct {
	msg string
}

func (m model) inti() tea.Cmd {
	return nil
}

func main() {
	fmt.Println("Welcome to Intuition")
}
