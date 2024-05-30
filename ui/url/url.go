package url

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type Model struct {
	style     lipgloss.Style
	textInput textinput.Model
	method    string
}

func New(widthInPercent float32) Model {
	return Model{
		style:     newDefaultStyle(widthInPercent),
		textInput: newTextInput(),
		method:    "GET",
	}
}

func newDefaultStyle(widthInPercent float32) lipgloss.Style {
	terminalWidth, _, _ := term.GetSize(os.Stdout.Fd())
	width := int(float32(terminalWidth) * widthInPercent)

	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Left).
		Width(width)
}

func newTextInput() textinput.Model {
	textInput := textinput.New()
	textInput.Placeholder = "Type your URL here..."
	textInput.Focus()
	textInput.CharLimit = 60
	return textInput
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textarea.Blink)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	}
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	terminalWidth, _, _ := term.GetSize(os.Stdout.Fd())
	width := m.style.GetWidth()

	oneSidePadding := (terminalWidth - width) / 2

	return lipgloss.JoinHorizontal(lipgloss.Bottom,
		strings.Repeat(" ", oneSidePadding),
		m.style.Render(m.method, m.textInput.View()),
		strings.Repeat(" ", oneSidePadding),
	)
}

func caculateWidth(args ...string) int {
	totalWidth := 0
	for _, arg := range args {
		totalWidth += len(arg)
	}

	return totalWidth
}
