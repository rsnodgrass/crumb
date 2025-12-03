package components

import "strings"

type Tags struct {
	tags []string
}

func NewTags() *Tags {
	return &Tags{
		tags: make([]string, 0),
	}
}

func (t *Tags) Add(tag string) {
	tag = strings.TrimSpace(tag)
	if tag != "" && !t.Contains(tag) {
		t.tags = append(t.tags, tag)
	}
}

func (t *Tags) Remove(index int) {
	if index >= 0 && index < len(t.tags) {
		t.tags = append(t.tags[:index], t.tags[index+1:]...)
	}
}

func (t *Tags) Contains(tag string) bool {
	for _, existingTag := range t.tags {
		if existingTag == tag {
			return true
		}
	}
	return false
}

func (t *Tags) All() []string {
	return t.tags
}

func (t *Tags) View() string {
	// Future: render as styled pills with lipgloss
	if len(t.tags) == 0 {
		return "Tags: (none)"
	}
	return "Tags: " + strings.Join(t.tags, ", ")
}
