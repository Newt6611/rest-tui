package panel

import (
	"os"
	"strings"

	"github.com/Newt6611/rest-tui/ui/key"
	"github.com/Newt6611/rest-tui/ui/response"
	"github.com/Newt6611/rest-tui/ui/url"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/term"
)

type PanelModel struct {
	title         string
	urlModel      url.Model
	responseModel response.Model
}

func New() *PanelModel {
	urlModel := url.New(0.5)
	responseModel := response.New(0.5, 0.5)

	return &PanelModel{
		title:         "REST UI",
		urlModel:      urlModel,
		responseModel: responseModel,
	}
}

func (m PanelModel) Init() tea.Cmd {
	return nil
}

func (m PanelModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case key.CtrlC, key.Q:
			return m, tea.Quit
		}
	}

	mo, cmd := m.urlModel.Update(msg)
	m.urlModel = mo.(url.Model)
	cmds = append(cmds, cmd)

	mo, cmd = m.responseModel.Update(msg)
	m.responseModel = mo.(response.Model)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m PanelModel) View() string {
	terminalWidth, _, _ := term.GetSize(os.Stdout.Fd())
	width := len(m.title)
	oneSidePadding := (terminalWidth - width) / 2

	return lipgloss.JoinVertical(lipgloss.Center,
		strings.Repeat(" ", oneSidePadding),
		m.title,
		strings.Repeat(" ", oneSidePadding),
		m.urlModel.View(),
		m.responseModel.View(),
	)
}
