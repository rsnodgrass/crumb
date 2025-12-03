package components

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	normalBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#313244")) // Surface

	focusedBorderStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#fab387")) // Peach

	labelStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6c7086")). // Overlay
			MarginBottom(1)

	placeholderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#a6adc8")) // Subtext
)

type Input struct {
	value       string
	placeholder string
	focused     bool
	width       int
}

func NewInput(placeholder string, width int) *Input {
	return &Input{
		placeholder: placeholder,
		width:       width,
		focused:     false,
	}
}

func (i *Input) SetValue(v string) {
	i.value = v
}

func (i *Input) Value() string {
	return i.value
}

func (i *Input) Focus() {
	i.focused = true
}

func (i *Input) Blur() {
	i.focused = false
}

func (i *Input) View() string {
	// Future: render with bubbles textinput or custom lipgloss
	if i.focused {
		return "[" + i.value + "_]"
	}
	return "[" + i.value + "]"
}

type TextArea struct {
	textarea textarea.Model
	label    string
	focused  bool
}

func NewTextArea(label, placeholder string) TextArea {
	ta := textarea.New()
	ta.Placeholder = placeholder
	ta.ShowLineNumbers = false
	ta.CharLimit = 0 // no limit by default

	// style the textarea
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.FocusedStyle.Base = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4")) // Text
	ta.BlurredStyle.Base = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4")) // Text
	ta.FocusedStyle.Placeholder = placeholderStyle
	ta.BlurredStyle.Placeholder = placeholderStyle

	return TextArea{
		textarea: ta,
		label:    label,
		focused:  false,
	}
}

func (t *TextArea) Focus() {
	t.focused = true
	t.textarea.Focus()
}

func (t *TextArea) Blur() {
	t.focused = false
	t.textarea.Blur()
}

func (t *TextArea) SetValue(s string) {
	t.textarea.SetValue(s)
}

func (t *TextArea) Value() string {
	return t.textarea.Value()
}

func (t TextArea) Update(msg tea.Msg) (TextArea, tea.Cmd) {
	var cmd tea.Cmd
	t.textarea, cmd = t.textarea.Update(msg)
	return t, cmd
}

func (t TextArea) View() string {
	// render label if present
	var labelView string
	if t.label != "" {
		labelView = labelStyle.Render(t.label) + "\n"
	}

	// choose border style based on focus
	borderStyle := normalBorderStyle
	if t.focused {
		borderStyle = focusedBorderStyle
	}

	// render textarea with border
	textareaView := borderStyle.Render(t.textarea.View())

	return labelView + textareaView
}
