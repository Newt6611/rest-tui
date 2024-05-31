package response

import (
	"os"
	"strings"

	"github.com/Newt6611/rest-tui/ui"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

const (
	Name string = "ResponsePanel"
)

type Model struct {
	focused        bool
	width          int
	height         int
	showHelp       bool
	responseString string
}

func New(widthInPercent, heightInPercent float32) Model {
	terminalWidth, terminalHeight, _ := term.GetSize(os.Stdout.Fd())
	width := int(float32(terminalWidth) * widthInPercent)
	height := int(float32(terminalHeight) * heightInPercent)

	return Model{
		width:    width,
		height:   height,
		showHelp: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}
	return m, cmd
}

func (m Model) View() string {
	terminalWidth, _ := ui.GetTerminalSize()
	oneSidePadding := (terminalWidth - m.width) / 2

	m.responseString = "Hi, this is a rest tui tool."
	styledOutput := ui.GetStyle(m.focused).
		Width(m.width).
		Height(m.height).
		Render(m.responseString)

	return lipgloss.JoinHorizontal(lipgloss.Bottom,
		strings.Repeat(" ", oneSidePadding),
		styledOutput,
		strings.Repeat(" ", oneSidePadding),
	)
}

func (m Model) ShowHelpPanel() bool {
	return m.showHelp
}

func (m Model) SetFocuse(b bool) ui.Model {
	m.focused = b
	return m
}

func (m Model) GetHelpKeyMap() help.KeyMap {
	return helpKeys
}
