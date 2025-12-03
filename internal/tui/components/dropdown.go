package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dropdown struct {
	options  []string
	selected int
	open     bool
	focused  bool
}

func NewDropdown(options []string, selectedIdx int) Dropdown {
	// ensure selected index is valid
	if selectedIdx < 0 || selectedIdx >= len(options) {
		selectedIdx = 0
	}

	return Dropdown{
		options:  options,
		selected: selectedIdx,
		open:     false,
		focused:  false,
	}
}

func (d Dropdown) Update(msg tea.Msg) (Dropdown, tea.Cmd) {
	if !d.focused {
		return d, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter", " ":
			d.open = !d.open
		case "up", "k":
			if d.open && d.selected > 0 {
				d.selected--
			}
		case "down", "j":
			if d.open && d.selected < len(d.options)-1 {
				d.selected++
			}
		}
	}

	return d, nil
}

func (d Dropdown) View() string {
	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#89dceb")).
		Bold(true)

	optionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4"))

	if d.open {
		return selectedStyle.Render("[" + d.Selected() + " ▼]")
	}
	return optionStyle.Render("[" + d.Selected() + "]")
}

func (d Dropdown) Selected() string {
	if d.selected >= 0 && d.selected < len(d.options) {
		return d.options[d.selected]
	}
	return ""
}

func (d *Dropdown) Focus() {
	d.focused = true
}

func (d *Dropdown) Blur() {
	d.focused = false
}
