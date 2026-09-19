package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	newFile                *os.File
}

var (
	vault       string
	cursorColor = lipgloss.Color("205")
)

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home directory", err)
	}

	vault = filepath.Join(homeDir, ".intuition")
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl+c", "ctrl+q":
			return m, tea.Quit

		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil

		case "enter":
			filename := m.newFileInput.Value()
			if filename != "" {
				filepath := fmt.Sprintf("%s/%s.md", vault, filename)

				if _, err := os.Stat(filepath); err == nil {
					return m, nil
				}

				newFile, err := os.Create(filepath)
				if err != nil {
					log.Fatal("Error creating file", err)
				}

				m.newFile = newFile
			}

			return m, nil
		}
	}
	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}

	return m, cmd
}

func (m model) View() tea.View {

	var style = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).Padding(0, 2, 0, 2)

	welcomemsg := style.Render("Welcome to Intuition")
	help := "Ctrl+N: new file - Ctrl+L: list - Esc: back/save - Ctrl+S: save - Ctrl+Q: quit"

	view := ""

	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}

	return tea.View{Content: fmt.Sprintf("\n%s\n\n%s\n\n%s", welcomemsg, view, help)}
}

func initializeMode() model {
	err := os.MkdirAll(vault, 0750)
	if err != nil {
		log.Fatal("Error in creating directory", err)
	}

	ti := textinput.New()
	ti.Placeholder = "Enter file name..."
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 25
	ti.SetVirtualCursor(true)
	ti.SetWidth(20)

	s := ti.Styles()
	s.Cursor.Color = cursorColor

	ti.SetStyles(s)

	return model{newFileInput: ti, createFileInputVisible: false}
}

func main() {
	p := tea.NewProgram(initializeMode())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Error occured: %v", err)
		os.Exit(1)
	}
}
