package response

import (
	"github.com/Newt6611/rest-tui/ui"
	"github.com/Newt6611/rest-tui/ui/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type shortcutModel struct {
	table          table.Model
	visible        bool
}

func newShortcutModel() shortcutModel {
	return shortcutModel{
		table:   newMethodTable(),
		visible: false,
	}
}

func newMethodTable() table.Model {
	columns := []table.Column{ }
	rows := []table.Row{ }

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(len(rows)),
		table.WithFocused(true),
	)

	// table styles
	s := table.DefaultStyles()
	s.Header = s.Header.Bold(true)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)
	return t
}

func (m shortcutModel) Init() tea.Cmd {
	return nil
}

func (m shortcutModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case key.Quit:
			m.visible = false

		case key.J:
			m.table.MoveDown(1)
		case key.K:
			m.table.MoveUp(1)
		}
	}

	return m, nil
}

func (m shortcutModel) View() string {
	terminalWidth, terminalHeight := ui.GetTerminalSize()

	return lipgloss.Place(terminalWidth,
		terminalHeight,
		lipgloss.Center,
		lipgloss.Center,
		ui.GetStyle(false).Render(m.table.View()),
	)
}
