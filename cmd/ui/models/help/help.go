package help

import (
	"github.com/charmbracelet/bubbles/key"
)

type KeyMap struct {
	Up    	key.Binding
	Down  	key.Binding
	Enter 	key.Binding
	Confirm key.Binding
	Help  	key.Binding
	Quit  	key.Binding
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down},
		{k.Enter, k.Confirm},
		{k.Help, k.Quit},
	}
}

var Keys = KeyMap {
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("Enter", " "),
		key.WithHelp("Enter/Space", "choose"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("y"),
		key.WithHelp("y", "confirm"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}
