package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type TagInput struct {
	tags          []string
	currentInput  string
	favoriteTags  []string
	focused       bool
}

func NewTagInput(favoriteTags []string) TagInput {
	return TagInput{
		tags:         make([]string, 0),
		currentInput: "",
		favoriteTags: favoriteTags,
		focused:      false,
	}
}

func (t TagInput) Update(msg tea.Msg) (TagInput, tea.Cmd) {
	if !t.focused {
		return t, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if t.currentInput != "" {
				tag := strings.TrimSpace(t.currentInput)
				if tag != "" && !t.contains(tag) {
					t.tags = append(t.tags, tag)
				}
				t.currentInput = ""
			}
		case "backspace":
			if t.currentInput == "" && len(t.tags) > 0 {
				t.tags = t.tags[:len(t.tags)-1]
			} else if len(t.currentInput) > 0 {
				t.currentInput = t.currentInput[:len(t.currentInput)-1]
			}
		default:
			if len(msg.String()) == 1 {
				t.currentInput += msg.String()
			}
		}
	}

	return t, nil
}

func (t TagInput) View() string {
	var parts []string

	tagStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#89dceb")).
		Background(lipgloss.Color("#313244")).
		Padding(0, 1).
		MarginRight(1)

	for _, tag := range t.tags {
		parts = append(parts, tagStyle.Render(tag))
	}

	if t.focused {
		inputStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#cdd6f4")).
			Background(lipgloss.Color("#1e1e2e")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#f9e2af")).
			Padding(0, 1)

		parts = append(parts, inputStyle.Render(t.currentInput+"_"))
	}

	if len(parts) == 0 {
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6c7086")).
			Render("(press enter to add tags)")
	}

	return strings.Join(parts, "")
}

func (t TagInput) Tags() []string {
	return t.tags
}

func (t *TagInput) Focus() {
	t.focused = true
}

func (t *TagInput) Blur() {
	t.focused = false
}

func (t *TagInput) Clear() {
	t.tags = make([]string, 0)
	t.currentInput = ""
}

func (t TagInput) contains(tag string) bool {
	for _, existingTag := range t.tags {
		if existingTag == tag {
			return true
		}
	}
	return false
}
