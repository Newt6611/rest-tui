package response

import (
	"github.com/Newt6611/rest-tui/ui/key"
	teaKey "github.com/charmbracelet/bubbles/key"
)

type helpKeyMap struct {
	Copy teaKey.Binding
}

func (k helpKeyMap) ShortHelp() []teaKey.Binding {
	return []teaKey.Binding{k.Copy}
}

func (k helpKeyMap) FullHelp() [][]teaKey.Binding {
	return [][]teaKey.Binding{
		{k.Copy},
	}
}

var helpKeys = helpKeyMap{
	Copy: teaKey.NewBinding(
		teaKey.WithKeys(key.C),
		teaKey.WithHelp(key.C, "copy to clipboard"),
	),
}
