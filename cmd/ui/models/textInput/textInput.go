package textInput

import (
	"PSU/cmd/ui/styles"
	"PSU/cmd/utils"
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

func (o *Output) GetOutput() string {
	return o.output
}

type model struct {
	textInput textinput.Model
	header string

	output *Output
	valid bool
	
	msg string
	err error
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
					input := m.textInput.Value()
					m.valid = !utils.ValidateProjectName(input)

					if !utils.ValidateProjectName(input) || len(input) == 0 {
						m.msg = "Invalid projetc name: '" + input + "'\n try not to use special symbols"

						return m, nil
					}
					
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