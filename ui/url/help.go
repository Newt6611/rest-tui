package url

import (
	"github.com/Newt6611/rest-tui/ui/key"
	teaKey "github.com/charmbracelet/bubbles/key"
)

type helpKeyMap struct {
	SelectMethod teaKey.Binding
	Paste        teaKey.Binding
}

func (k helpKeyMap) ShortHelp() []teaKey.Binding {
	return []teaKey.Binding{k.SelectMethod, k.Paste}
}

func (k helpKeyMap) FullHelp() [][]teaKey.Binding {
	return [][]teaKey.Binding{
		{k.SelectMethod},
		{k.Paste},
	}
}

var helpKeys = helpKeyMap{
	SelectMethod: teaKey.NewBinding(
		teaKey.WithKeys(key.Question),
		teaKey.WithHelp(key.Question, "select http method"),
	),
	Paste: teaKey.NewBinding(
		teaKey.WithKeys(key.CtrlP),
		teaKey.WithHelp(key.CtrlP, "paste"),
	),
}
