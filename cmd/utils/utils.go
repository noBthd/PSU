package utils

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func SplitLines(s string) []string {
	var lines []string
	current := ""
	for _, r := range s {
		if r == '\n' {
			lines = append(lines, current)
			current = ""
		} else {
			current += string(r)
		}
	}
	if current != "" {
		lines = append(lines, current)
	}
	return lines
}

func InterpolateColor(from, to lipgloss.Color, t float64) lipgloss.Color {
	fr, fg, fb := HexToRGB(from)
	tr, tg, tb := HexToRGB(to)

	r := uint8(float64(fr)*(1-t) + float64(tr)*t)
	g := uint8(float64(fg)*(1-t) + float64(tg)*t)
	b := uint8(float64(fb)*(1-t) + float64(tb)*t)

	return lipgloss.Color(fmt.Sprintf("#%02X%02X%02X", r, g, b))
}

func HexToRGB(c lipgloss.Color) (r, g, b uint8) {
	var rr, gg, bb int
	fmt.Sscanf(string(c), "#%02X%02X%02X", &rr, &gg, &bb)
	return uint8(rr), uint8(gg), uint8(bb)
}

func ValidateProjectName(n string) bool {
	nonValidChars := [11]string{" ", ".", "\\", "<", ">", ":", "\"", "/", "|", "?", "*"}
	
	for i := 0; i < len(nonValidChars); i++ {
		if strings.Contains(n, nonValidChars[i]) {
			return false
		}

		if len(n) == 0 {
			return false
		}
	}

	return true
}
