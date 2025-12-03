package tui

import (
	"os"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// toast styling with improved borders
var (
	successToastStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#a6e3a1")).
				Background(lipgloss.Color("#1e1e2e")).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#a6e3a1")).
				Padding(0, 2).
				Bold(true)

	errorToastStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f38ba8")).
			Background(lipgloss.Color("#1e1e2e")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#f38ba8")).
			Padding(0, 2).
			Bold(true)
)

// ToastHideMsg is sent after toast timeout
type ToastHideMsg struct{}

// RenderToast renders a centered toast notification at the bottom of the screen
func RenderToast(msg string, isError bool, width int) string {
	var style lipgloss.Style
	var icon string

	if isError {
		style = errorToastStyle
		icon = "✗"
	} else {
		style = successToastStyle
		icon = "✓"
	}

	// format message with icon
	content := icon + " " + msg

	// render with style
	styledContent := style.Render(content)

	// calculate rendered width to center properly
	renderedWidth := lipgloss.Width(styledContent)
	leftPadding := (width - renderedWidth) / 2
	if leftPadding < 0 {
		leftPadding = 0
	}

	// center the toast
	return strings.Repeat(" ", leftPadding) + styledContent
}

// HideToastAfter returns a command that sends ToastHideMsg after the specified duration
func HideToastAfter(duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return ToastHideMsg{}
	})
}

// FormatSaveError returns a user-friendly error message based on the error type
func FormatSaveError(err error) string {
	if os.IsPermission(err) {
		return "Permission denied - check directory permissions"
	}
	if os.IsNotExist(err) {
		return "Directory not found - run 'prompt-share init' first"
	}
	return err.Error()
}
