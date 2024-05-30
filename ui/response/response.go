package response

import (
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type Model struct {
	style lipgloss.Style
}

func New(widthInPercent, heightInPercent float32) Model {
	return Model{
		style: newDefaultStyle(widthInPercent, heightInPercent),
	}
}

func newDefaultStyle(widthInPercent, heightInPercent float32) lipgloss.Style {
	terminalWidth, terminalHeight, _ := term.GetSize(os.Stdout.Fd())
	width := int(float32(terminalWidth) * widthInPercent)
	height := int(float32(terminalHeight) * heightInPercent)

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Left).
		Width(width).
		Height(height)
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
	terminalWidth, _, _ := term.GetSize(os.Stdout.Fd())
	width := m.style.GetWidth()
	oneSidePadding := (terminalWidth - width) / 2

	return lipgloss.JoinHorizontal(lipgloss.Bottom,
		strings.Repeat(" ", oneSidePadding),
		m.style.Render("{\n	\"foor\":\"bar\"\n}"),
		strings.Repeat(" ", oneSidePadding),
	)
}
