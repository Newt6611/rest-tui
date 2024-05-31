package ui

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type PannelOrder int

const (
	UrlIndex      PannelOrder = 0
	ResponseIndex PannelOrder = 1
)

type Model interface {
	tea.Model

	ShowHelpPanel() bool
	SetFocuse(bool) Model
	GetHelpKeyMap() help.KeyMap
}

func GetStyle(focus bool) lipgloss.Style {
	if focus {
		return FocusedStyle()
	} else {
		return NormalStyle()
	}

}

func NormalStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#185900"))
}

func FocusedStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#fff72b"))
}
