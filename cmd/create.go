package cmd

import (
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"

	p "PSU/cmd/project"
	"PSU/cmd/ui/models/multiInput"
	textinput "PSU/cmd/ui/models/textInput"
	"PSU/cmd/utils"
)

// JUST LOGO
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

// VARS
var (
	ProjectNameIsValid = false
	tprogram 	*tea.Program
	Project 	*p.Project
)

// MAIN
func init() {
	rootCmd.AddCommand(createCmd)
}

type listOptions struct {
	options []string
}

type Options struct {
	project 		*p.Project

	ProjectName 	*textinput.Output
	ProjectLang 	*multiInput.Selection
	ProjectGit  	*multiInput.Selection
	ProjectGitLink	*textinput.Output

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
			project: 	 	&p.Project{},

			ProjectName: 	&textinput.Output{},
			ProjectLang: 	&multiInput.Selection{},
			ProjectGit:  	&multiInput.Selection{},
			ProjectGitLink: &textinput.Output{},
		}

		simpleChoose := listOptions {
			options: []string{
				"yes",
				"no",
			},
		}

		if !ProjectNameIsValid {
			tprogram = tea.NewProgram(textinput.InitialModel(options.ProjectName, "Enter project name:", "Name?"))
			if _, err := tprogram.Run(); err != nil {
				log.Printf("error in running the program")
				cobra.CheckErr(err)
			}
		}

		tprogram = tea.NewProgram(multiInput.InitialModel("Create git repo?:", options.ProjectGit, simpleChoose.options))
		if _, err := tprogram.Run(); err != nil {
			cobra.CheckErr(err)
		}

		log.Print(options.ProjectGit.Choice)
		log.Print(simpleChoose.options[0])

		if options.ProjectGit.Choice == simpleChoose.options[0] {
			tprogram = tea.NewProgram(textinput.InitialModel(options.ProjectGitLink, "Enter github link:", "Link?"))
			if _, err := tprogram.Run(); err != nil {
				log.Printf("error in running the program")
				cobra.CheckErr(err)
			}	
		}
	},
}