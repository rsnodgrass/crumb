package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	helpWidth  = 41
	helpHeight = 13
)

var (
	helpTitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#b4befe")).
			Bold(true).
			Align(lipgloss.Center).
			Width(helpWidth - 4)

	helpTextStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cdd6f4"))

	helpBoxStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#313244")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#b4befe")).
			Padding(0, 1).
			Width(helpWidth)
)

// RenderHelpOverlay renders a centered help overlay with keyboard shortcuts
func RenderHelpOverlay(width, height int) string {
	shortcuts := []struct {
		key  string
		desc string
	}{
		{"Tab", "Next field"},
		{"Shift+Tab", "Previous field"},
		{"Ctrl+S", "Save and exit"},
		{"Esc", "Cancel and exit"},
		{"/", "Open tool selector"},
		{"?", "Toggle this help"},
	}

	var content strings.Builder

	// title
	content.WriteString(helpTitleStyle.Render("Keyboard Shortcuts"))
	content.WriteString("\n\n")

	// shortcuts
	for _, shortcut := range shortcuts {
		keyStyle := helpTextStyle.Copy().Bold(true)
		key := keyStyle.Render(shortcut.key)
		desc := helpTextStyle.Render(shortcut.desc)

		// pad key to align descriptions
		keyWidth := 14
		paddedKey := key + strings.Repeat(" ", keyWidth-lipgloss.Width(key))

		content.WriteString("  " + paddedKey + desc + "\n")
	}

	content.WriteString("\n")
	content.WriteString(helpTextStyle.Copy().Italic(true).Render("         Press any key to close"))

	box := helpBoxStyle.Render(content.String())

	// center the overlay
	return lipgloss.Place(
		width,
		height,
		lipgloss.Center,
		lipgloss.Center,
		box,
	)
}
