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

	// FocusedBorderStyle - Peach color borders with padding
	FocusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(Peach)).
				Padding(0, 1)

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
			Foreground(lipgloss.Color(Red)).
			Italic(true)

	// ErrorFieldBorderStyle - Red border for invalid fields
	ErrorFieldBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(Red)).
				Padding(0, 1)

	// UnfocusedBorderStyle - Surface color borders (unfocused)
	UnfocusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(Surface)).
				Padding(0, 1)

	// HelpStyle - Subtext, dim
	HelpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Subtext)).
			Faint(true)

	// FooterStyle - Overlay color, italic
	FooterStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Overlay)).
			Italic(true)

	// FooterKeyStyle - Blue color for keys
	FooterKeyStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Blue)).
			Bold(true)

	// TitleStyle - App title (Lavender, larger)
	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Lavender)).
			Bold(true)

	// SubheaderStyle - Output directory hint (Subtext)
	SubheaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(Subtext))

	// FocusedLabelStyle - Focused field label (Peach)
	FocusedLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(Peach)).
				Bold(true)

	// UnfocusedLabelStyle - Unfocused field label (Overlay)
	UnfocusedLabelStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(Overlay))

	// SectionStyle - Section borders with padding
	SectionStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(Surface)).
			Padding(1, 2)

	// DropdownClosedStyle - closed dropdown with surface border
	DropdownClosedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(Text)).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(Surface)).
				Padding(0, 1)

	// DropdownOpenStyle - open dropdown with peach border
	DropdownOpenStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(Text)).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color(Peach)).
				Padding(0, 1)

	// DropdownSelectedStyle - selected option in peach
	DropdownSelectedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(Peach)).
				Bold(true)

	// DropdownDefaultStyle - default indicator in overlay
	DropdownDefaultStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(Overlay)).
				Faint(true)
)
