package textinput

import (
	"PSU/cmd/ui/styles"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type (
	errMsg error
)

type Output struct {
	output string
}

func (o *Output) Update(val string) {
	o.output = val
}

type model struct {
	textInput textinput.Model
	err error
	output *Output
	header string
}

func InitialModel(o *Output, header string) model {
	ti := textinput.New()
	ti.Placeholder = "language?"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return model{
		textInput: ti,
		err: nil,
		output: o,
		header: header,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

		switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n",
		styles.TitleStyle.Render(m.header),
		m.textInput.View(),
	)
}