package multiInput

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#31E981"))
	titleStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#4E148C"))
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
		header: titleStyle.Render(header),
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
	s := "\n" + m.header + "\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = focusedStyle.Render(">")
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = focusedStyle.Render("x")
		}

		s += fmt.Sprintf("\n%s [%s] %s\n", cursor, checked, choice)
	}

	return s
}

func (m model) Init() tea.Cmd {
	return nil
}