package styles

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	DefaultStyle		= lipgloss.NewStyle().
							Foreground(lipgloss.Color("#97DFFC"))

	FocusedStyle 		= lipgloss.NewStyle().
							Foreground(lipgloss.Color("#4E148C"))

	ChooseStyle 		= lipgloss.NewStyle().
							Foreground(lipgloss.Color("#4E148C")).
							Bold(true).
							Italic(true)

	TextFocusedStyle 	= lipgloss.NewStyle().
							Foreground(lipgloss.Color("#858AE3"))

	TitleStyle 			= lipgloss.NewStyle().
							Background(lipgloss.Color("#4E148C"))
)