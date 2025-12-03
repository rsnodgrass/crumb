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
	suggestions   []string
	focused       bool
}

func NewTagInput(favoriteTags []string) TagInput {
	ti := TagInput{
		tags:         make([]string, 0),
		currentInput: "",
		favoriteTags: favoriteTags,
		suggestions:  make([]string, 0),
		focused:      false,
	}
	ti.updateSuggestions()
	return ti
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
				t.updateSuggestions()
			}
		case "backspace":
			if t.currentInput == "" && len(t.tags) > 0 {
				t.tags = t.tags[:len(t.tags)-1]
				t.updateSuggestions()
			} else if len(t.currentInput) > 0 {
				t.currentInput = t.currentInput[:len(t.currentInput)-1]
				t.updateSuggestions()
			}
		case "1", "2", "3", "4", "5":
			// quick select suggestion by number
			idx := int(msg.String()[0] - '1')
			if idx < len(t.suggestions) {
				tag := t.suggestions[idx]
				if !t.contains(tag) {
					t.tags = append(t.tags, tag)
					t.currentInput = ""
					t.updateSuggestions()
				}
			}
		default:
			if len(msg.String()) == 1 {
				t.currentInput += msg.String()
				t.updateSuggestions()
			}
		}
	}

	return t, nil
}

func (t TagInput) View() string {
	var b strings.Builder

	// styles
	tagStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#1e1e2e")).
		Background(lipgloss.Color("#89b4fa")).
		Padding(0, 1).
		MarginRight(1)

	inputPlaceholderStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6c7086")).
		Faint(true)

	suggestionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6c7086")).
		Italic(true)

	suggestionKeyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#89b4fa")).
		Bold(true)

	inputStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4")).
		Background(lipgloss.Color("#1e1e2e")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#f9e2af")).
		Padding(0, 1)

	// render selected tags as badges
	for _, tag := range t.tags {
		b.WriteString(tagStyle.Render(tag))
		b.WriteString(" ")
	}

	// render input field or placeholder
	if t.focused {
		if t.currentInput == "" {
			b.WriteString(inputPlaceholderStyle.Render("[+add tag]"))
		} else {
			b.WriteString(inputStyle.Render(t.currentInput + "_"))
		}

		// render suggestions below
		if len(t.suggestions) > 0 {
			b.WriteString("\n")
			b.WriteString(suggestionStyle.Render("  Suggestions: "))

			maxSuggestions := min(5, len(t.suggestions))
			for i := 0; i < maxSuggestions; i++ {
				if i > 0 {
					b.WriteString(" • ")
				}
				// show number key hint
				b.WriteString(suggestionKeyStyle.Render(string(rune('1' + i))))
				b.WriteString(":")
				b.WriteString(t.suggestions[i])
			}
		}
	} else if len(t.tags) == 0 {
		b.WriteString(inputPlaceholderStyle.Render("(press enter to add tags)"))
	}

	return b.String()
}

func (t TagInput) Tags() []string {
	return t.tags
}

func (t *TagInput) Focus() {
	t.focused = true
	t.updateSuggestions()
}

func (t *TagInput) Blur() {
	t.focused = false
}

func (t *TagInput) Clear() {
	t.tags = make([]string, 0)
	t.currentInput = ""
	t.updateSuggestions()
}

func (t TagInput) contains(tag string) bool {
	for _, existingTag := range t.tags {
		if existingTag == tag {
			return true
		}
	}
	return false
}

func (t *TagInput) updateSuggestions() {
	t.suggestions = make([]string, 0)
	input := strings.ToLower(t.currentInput)

	// filter favorites based on input and exclude already selected tags
	for _, fav := range t.favoriteTags {
		if t.contains(fav) {
			continue
		}

		if t.currentInput == "" || strings.Contains(strings.ToLower(fav), input) {
			t.suggestions = append(t.suggestions, fav)
		}
	}
}
