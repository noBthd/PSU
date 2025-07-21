package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	"PSU/cmd/ui/multiInput"
	"PSU/cmd/utils"
)

const logo = `
 ███████████   █████████  █████  █████
░░███░░░░░███ ███░░░░░███░░███  ░░███
 ░███    ░███░███    ░░░  ░███   ░███
 ░██████████ ░░█████████  ░███   ░███
 ░███░░░░░░   ░░░░░░░░███ ░███   ░███
 ░███         ███    ░███ ░███   ░███
 █████       ░░█████████  ░░████████
░░░░░         ░░░░░░░░░    ░░░░░░░░
`

// var (
// 	logoStyle 			= lipgloss.NewStyle().
// 							Foreground(lipgloss.Color("#31E981")).
// 							Bold(true).
// 							Align(lipgloss.Center)
// )

func init() {
	rootCmd.AddCommand(createCmd)
}

type listOptions struct {
	options []string
}

type Options struct {
	ProjectName *string
	ProjectLang *multiInput.Selection
}

var createCmd = &cobra.Command{
	Use: "run",
	Short: "Short description",
	Long: ".",

	Run: func(cmd *cobra.Command, args []string) {
		lines := []string{}
		lines = append(lines, utils.SplitLines(logo)...)

		colors := []lipgloss.Color{
			lipgloss.Color("#2C0735"), 
			lipgloss.Color("#613DC1"), 
			lipgloss.Color("#97DFFC"), 
		}

		gradientLines := make([]string, len(lines))
		n := len(lines)

		for i, line := range lines {
			var color lipgloss.Color
			ratio := float64(i) / float64(n-1)

			switch {
			case ratio < 0.5:
				color = utils.InterpolateColor(colors[0], colors[1], ratio*2)
			default:
				color = utils.InterpolateColor(colors[1], colors[2], (ratio-0.5)*2)
			}

			style := lipgloss.NewStyle().
				Foreground(color).
				Bold(true).
				Align(lipgloss.Center)

			gradientLines[i] = style.Render(line)
		}

		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, gradientLines...) + "\n")

		options := Options {
			ProjectLang: &multiInput.Selection{},
		}

		listOfStuff := listOptions {
			options: []string{
				"C++",
				"Golang",
			},
		}

		tprogram := tea.NewProgram(multiInput.InitialModel("Choose language to use:", options.ProjectLang, listOfStuff.options))
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}
	},
}