package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Catppuccin Mocha color palette
const (
	Lavender = "#b4befe" // header/title
	Overlay  = "#6c7086" // labels
	Surface  = "#313244" // borders
	Mauve    = "#cba6f7" // tool selector accent
	Blue     = "#89b4fa" // tags
	Green    = "#a6e3a1" // success
	Red      = "#f38ba8" // error
	Text     = "#cdd6f4" // prompt text
	Peach    = "#fab387" // cursor/focus
	Base     = "#1e1e2e" // background
	Subtext  = "#a6adc8" // secondary text
)

var (
	// HeaderStyle - Lavender, bold
	HeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Lavender)).
			Bold(true)

	// LabelStyle - Overlay color
	LabelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Overlay))

	// BorderStyle - Surface color borders
	BorderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(Surface))

	// FocusedBorderStyle - Peach color borders
	FocusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(Peach))

	// InputStyle - Text color
	InputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Text))

	// TagStyle - Blue background, rounded
	TagStyle = lipgloss.NewStyle().
			Background(lipgloss.Color(Blue)).
			Foreground(lipgloss.Color(Base)).
			Padding(0, 1).
			MarginRight(1).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(Blue))

	// SuccessStyle - Green
	SuccessStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Green))

	// ErrorStyle - Red
	ErrorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Red))

	// HelpStyle - Subtext, dim
	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Subtext)).
			Faint(true)
)
