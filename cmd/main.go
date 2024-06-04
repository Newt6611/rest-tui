package main

import (
	"fmt"
	"os"
	"github.com/Newt6611/rest-tui/ui/panel"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	//
	//2
	mainWindow := panel.New()
	p := tea.NewProgram(*mainWindow, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
