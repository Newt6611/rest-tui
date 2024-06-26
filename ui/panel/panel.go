package panel

import (
	"os"

	"github.com/Newt6611/rest-tui/ui"
	"github.com/Newt6611/rest-tui/ui/key"
	"github.com/Newt6611/rest-tui/ui/response"
	"github.com/Newt6611/rest-tui/ui/url"
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type Model struct {
	helpView     help.Model
	title        string
	subModels    []ui.Model
	focusedIndex int // Current Focused Panel Index
}

func New() *Model {
	urlModel := url.New(0.5)
	{
		m := urlModel.SetFocuse(true)
		urlModel = m.(url.Model)
	}
	responseModel := response.New(0.5, 0.5)

	subModels := []ui.Model{
		urlModel,
		responseModel,
	}

	return &Model{
		title:        "REST UI",
		subModels:    subModels,
		focusedIndex: 0,
		helpView: help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.helpView.Width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case key.CtrlC, key.Esc:
			return m, tea.Quit
		case key.Quit:
			if m.focusedIndex != int(ui.UrlIndex) {
				return m, tea.Quit
			}
		case key.Tab:
			model := m.subModels[m.focusedIndex]
			if !model.ShowHelpPanel() {
				m.updateFocusedModel()
			}
		}
	}

	mo, cmd := m.subModels[m.focusedIndex].Update(msg)
	m.subModels[m.focusedIndex] = mo.(ui.Model)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	terminalWidth, terminalHeight, _ := term.GetSize(os.Stdout.Fd())

	model := m.subModels[m.focusedIndex]
	if model.ShowHelpPanel() {
		return model.View()
	} else {
		styledOutput := lipgloss.JoinVertical(lipgloss.Center,
			m.title,
			m.subModels[ui.UrlIndex].View(),
			m.subModels[ui.ResponseIndex].View(),
			m.helpView.View(m.subModels[m.focusedIndex].GetHelpKeyMap()),
		)
		return lipgloss.Place(terminalWidth, terminalHeight, lipgloss.Center, lipgloss.Center, styledOutput)
	}
}

func (m *Model) updateFocusedModel() {
	currentModel := m.subModels[m.focusedIndex]
	m.subModels[m.focusedIndex] = currentModel.SetFocuse(false)

	m.focusedIndex++
	if m.focusedIndex >= len(m.subModels) {
		m.focusedIndex = 0
	}
	nextModel := m.subModels[m.focusedIndex]

	m.subModels[m.focusedIndex] = nextModel.SetFocuse(true)
}
