package response

import (
	teaKey "github.com/charmbracelet/bubbles/key"
)

type helpKeyMap struct {
}

func (k helpKeyMap) ShortHelp() []teaKey.Binding {
	return []teaKey.Binding{}
}

func (k helpKeyMap) FullHelp() [][]teaKey.Binding {
	return [][]teaKey.Binding{
		{},
	}
}

var helpKeys = helpKeyMap {
	// SelectMethod: teaKey.NewBinding(
	// 	teaKey.WithKeys(key.Question),
	// 	teaKey.WithHelp(key.Question, "select http method"),
	// ),
}
