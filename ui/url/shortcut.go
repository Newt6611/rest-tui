package url

import (
	"net/http"
	"strings"

	"github.com/Newt6611/rest-tui/ui"
	"github.com/Newt6611/rest-tui/ui/key"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type shortcutModel struct {
	table          table.Model
	visible        bool
	methods        []string
	selectedMethod string
}

func newShortcutModel() shortcutModel {
	return shortcutModel{
		table:   newMethodTable(),
		visible: false,
		methods: []string{},
	}
}

func newMethodTable() table.Model {
	columns := []table.Column{
		{Title: "ShortCut", Width: 10},
		{Title: "Method", Width: 15},
	}
	rows := []table.Row{
		{key.Get, http.MethodGet},
		{key.Post, http.MethodPost},
		{key.Put, http.MethodPut},
		{key.Delete, http.MethodDelete},
		{key.Head, http.MethodHead},
		{key.Patch, http.MethodPatch},
		{key.Connect, http.MethodConnect},
		{key.Options, http.MethodOptions},
		{key.Trace, http.MethodTrace},
	}

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

		case key.Enter:
			row := m.table.SelectedRow()
			method := row[1]
			m.selectedMethod = method
			m.visible = false

		case key.Get:
			m.selectedMethod = http.MethodGet
			m.table.SetCursor(0)
			m.visible = false
		case key.Post:
			m.selectedMethod = http.MethodPost
			m.table.SetCursor(1)
			m.visible = false
		case key.Put:
			m.selectedMethod = http.MethodPut
			m.table.SetCursor(2)
			m.visible = false
		case key.Delete:
			m.selectedMethod = http.MethodDelete
			m.table.SetCursor(3)
			m.visible = false
		case key.Head:
			m.selectedMethod = http.MethodHead
			m.table.SetCursor(4)
			m.visible = false
		case key.Patch:
			m.selectedMethod = http.MethodPatch
			m.table.SetCursor(5)
			m.visible = false
		case key.Connect:
			m.selectedMethod = http.MethodConnect
			m.table.SetCursor(6)
			m.visible = false
		case key.Options:
			m.selectedMethod = http.MethodOptions
			m.table.SetCursor(7)
			m.visible = false
		case key.Trace:
			m.selectedMethod = http.MethodTrace
			m.table.SetCursor(8)
			m.visible = false
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

func (m shortcutModel) getSelectedMethod() string {
	return m.selectedMethod
}

// get max width from list of method,
// for example:
// GET, POST, DELETE
// will return 6 len(DELETE)
func (m shortcutModel) getMaxWidth() int {
	width := 0
	for _, m := range m.methods {
		if len(m) > width {
			width = len(m)
		}
	}
	return width
}

func (m shortcutModel) getMaxHeight() int {
	return len(m.methods)
}

func (m shortcutModel) getMethodString() string {
	sb := strings.Builder{}
	for _, s := range m.methods {
		sb.WriteString(s)
		sb.WriteString("\n")
	}
	return sb.String()
}
