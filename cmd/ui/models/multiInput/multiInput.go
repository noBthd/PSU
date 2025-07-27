package multiInput

import (
	"PSU/cmd/ui/styles"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Selection struct {
	choice string
}

func (s *Selection) Update(val string) {
	s.choice = val
}

type model struct {
	choices []string
	cursor int
	selected map[int]struct{}
	choice *Selection
	header string
}

func InitialModel(header string, selection *Selection, choices []string) model {
	return model {
		header: styles.TitleStyle.Render(header),
		choices: choices,
		selected: make(map[int]struct{}),
		choice: selection,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

			case "ctrl+c", "q":
				return m, tea.Quit

			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}

			case "down", "j":
				if m.cursor < len(m.choices)-1 {
					m.cursor++
				}

			case "enter", " ":
				if len(m.selected) == 1 {
					m.selected = make(map[int]struct{})
				}

				_, ok := m.selected[m.cursor]
				if ok {
					delete(m.selected, m.cursor)
				} else {
					m.selected[m.cursor] = struct{}{}
				}

		}
	}

	return m, nil
}

func (m model) View() string {
	s := m.header + "\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = styles.FocusedStyle.Render(">")
			choice = styles.TextFocusedStyle.Render(choice)
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = styles.ChooseStyle.Render("x")
			choice = styles.ChooseStyle.Render(choice)
		}

		s += fmt.Sprintf("\n%s [%s] %s\n", cursor, checked, styles.DefaultStyle.Render(choice))
	}

	return s
}

func (m model) Init() tea.Cmd {
	return nil
}