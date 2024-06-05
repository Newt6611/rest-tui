package response

import (
	"os"

	"github.com/Newt6611/rest-tui/ui"
	"github.com/Newt6611/rest-tui/ui/key"
	"github.com/atotto/clipboard"
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
	bgColor        lipgloss.TerminalColor
}

func New(widthInPercent, heightInPercent float32) Model {
	terminalWidth, terminalHeight, _ := term.GetSize(os.Stdout.Fd())
	width := int(float32(terminalWidth) * widthInPercent)
	height := int(float32(terminalHeight) * heightInPercent)

	return Model{
		width:          width,
		height:         height,
		showHelp:       false,
		responseString: "Hi, this is a rest tui tool.",
	}
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
		case key.C:
			clipboard.WriteAll(m.responseString)
			m.bgColor = lipgloss.Color("#FFF")

			cmd = ui.StartCopyEffect(ui.CopyEffectTime)
			cmds = append(cmds, cmd)
		}

	case ui.CopyEffect:
		// When true, means timer end, then we can set background color back to default
		b := bool(msg)
		if b {
			m.bgColor = ui.GetStyle(m.focused).GetBackground()
		}

	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	styledOutput := ui.GetStyle(m.focused).
		Width(m.width).
		Height(m.height).
		Background(m.bgColor).
		Render(m.responseString)

	return styledOutput
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
