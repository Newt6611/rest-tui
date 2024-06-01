package url

import (
	"os"
	"strings"

	"github.com/Newt6611/rest-tui/ui"
	"github.com/Newt6611/rest-tui/ui/key"
	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

const Name string = "UrlPanel"

type Model struct {
	shortcutModel shortcutModel
	textInput     textinput.Model

	width   int
	method  string // Http method
	focused bool
}

func New(widthInPercent float32) Model {
	terminalWidth, _, _ := term.GetSize(os.Stdout.Fd())
	width := int(float32(terminalWidth) * widthInPercent)

	model := Model{
		width:     width,
		textInput: newTextInput(),
		method:    "GET",
	}

	model.shortcutModel = newShortcutModel()
	return model
}

func newTextInput() textinput.Model {
	textInput := textinput.New()
	textInput.Placeholder = "Type your URL here..."
	textInput.Focus()
	textInput.CharLimit = 60
	return textInput
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case key.Question:
			m.shortcutModel.visible = true
		// case key.CtrlP: // Paste from clipboard
		// 	clipboardText, _ := clipboard.ReadAll()
		// 	preText := m.textInput.Value()
		// 	newText := preText + clipboardText
		// 	m.textInput.SetValue(newText)
		// 	m.textInput.SetCursor(len(newText))
		// case key.CtrlShiftC: // Copy from the url text input
		// 	text := m.textInput.Value()
		// 	clipboard.WriteAll(text)
		}
	}

	if m.focused && !m.ShowHelpPanel() {
		m.textInput, cmd = m.textInput.Update(msg)
		cmds = append(cmds, cmd)
	}

	if m.ShowHelpPanel() {
		var mm tea.Model
		mm, cmd = m.shortcutModel.Update(msg)
		cmds = append(cmds, cmd)
		m.shortcutModel = mm.(shortcutModel)
		// Update current selected method from shortcut model
		m.method = m.shortcutModel.getSelectedMethod()
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	terminalWidth, _ := ui.GetTerminalSize()

	oneSidePadding := (terminalWidth - m.width) / 2

	styledOutput := ui.GetStyle(m.focused).
		Width(m.width).
		Render(m.method, m.textInput.View())

	if !m.shortcutModel.visible {
		return lipgloss.JoinHorizontal(lipgloss.Bottom,
			strings.Repeat(" ", oneSidePadding),
			styledOutput,
			strings.Repeat(" ", oneSidePadding))
	} else {
		return m.shortcutModel.View()
	}
}

func (m Model) ShowHelpPanel() bool {
	return m.shortcutModel.visible
}

func (m Model) SetFocuse(b bool) ui.Model {
	m.focused = b
	if m.focused {
		m.textInput.Cursor.SetMode(cursor.CursorBlink)
	} else {
		m.textInput.Cursor.SetMode(cursor.CursorHide)
	}
	return m
}

func (m Model) GetHelpKeyMap() help.KeyMap {
	return helpKeys
}

func caculateWidth(args ...string) int {
	totalWidth := 0
	for _, arg := range args {
		totalWidth += len(arg)
	}

	return totalWidth
}
