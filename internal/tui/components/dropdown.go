package components

import (
	"strings"
	"unicode"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Dropdown struct {
	options     []string
	selected    int
	open        bool
	focused     bool
	defaultTool string
	filter      string
	filtered    []int // indices of matching options
}

func NewDropdown(options []string, selectedIdx int, defaultTool string) Dropdown {
	// ensure selected index is valid
	if selectedIdx < 0 || selectedIdx >= len(options) {
		selectedIdx = 0
	}

	return Dropdown{
		options:     options,
		selected:    selectedIdx,
		open:        false,
		focused:     false,
		defaultTool: defaultTool,
		filter:      "",
		filtered:    []int{},
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
			// clear filter when closing
			if !d.open {
				d.filter = ""
				d.filtered = []int{}
			}
		case "up", "k":
			if d.open {
				d.movePrev()
			}
		case "down", "j":
			if d.open {
				d.moveNext()
			}
		case "backspace":
			if d.open && len(d.filter) > 0 {
				d.filter = d.filter[:len(d.filter)-1]
				d.updateFiltered()
			}
		case "esc":
			if d.open {
				d.open = false
				d.filter = ""
				d.filtered = []int{}
			}
		default:
			// handle type-to-filter when dropdown is open
			if d.open {
				char := msg.String()
				if len(char) == 1 && (unicode.IsLetter(rune(char[0])) || unicode.IsDigit(rune(char[0]))) {
					d.filter += strings.ToLower(char)
					d.updateFiltered()
				}
			}
		}
	}

	return d, nil
}

// updateFiltered rebuilds the filtered indices based on current filter
func (d *Dropdown) updateFiltered() {
	d.filtered = []int{}
	if d.filter == "" {
		return
	}

	filter := strings.ToLower(d.filter)
	for i, opt := range d.options {
		if strings.Contains(strings.ToLower(opt), filter) {
			d.filtered = append(d.filtered, i)
		}
	}

	// adjust selection to first filtered option if current selection is not in filtered list
	if len(d.filtered) > 0 {
		inFiltered := false
		for _, idx := range d.filtered {
			if idx == d.selected {
				inFiltered = true
				break
			}
		}
		if !inFiltered {
			d.selected = d.filtered[0]
		}
	}
}

// movePrev moves selection to previous option (filtered or unfiltered)
func (d *Dropdown) movePrev() {
	options := d.getVisibleOptions()
	if len(options) == 0 {
		return
	}

	// find current position in visible options
	currentPos := -1
	for i, idx := range options {
		if idx == d.selected {
			currentPos = i
			break
		}
	}

	if currentPos > 0 {
		d.selected = options[currentPos-1]
	}
}

// moveNext moves selection to next option (filtered or unfiltered)
func (d *Dropdown) moveNext() {
	options := d.getVisibleOptions()
	if len(options) == 0 {
		return
	}

	// find current position in visible options
	currentPos := -1
	for i, idx := range options {
		if idx == d.selected {
			currentPos = i
			break
		}
	}

	if currentPos >= 0 && currentPos < len(options)-1 {
		d.selected = options[currentPos+1]
	}
}

// getVisibleOptions returns indices of options to display (filtered or all)
func (d Dropdown) getVisibleOptions() []int {
	if len(d.filtered) > 0 {
		return d.filtered
	}

	// return all options
	result := make([]int, len(d.options))
	for i := range d.options {
		result[i] = i
	}
	return result
}

func (d Dropdown) View() string {
	if d.open {
		return d.viewOpen()
	}
	return d.viewClosed()
}

// viewClosed renders the closed dropdown with selected value
func (d Dropdown) viewClosed() string {
	closedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#313244")).
		Padding(0, 1)

	return closedStyle.Render(d.Selected() + " ▼")
}

// viewOpen renders the open dropdown with options list
func (d Dropdown) viewOpen() string {
	var b strings.Builder

	openStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4")).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#fab387")).
		Padding(0, 1)

	selectedStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#fab387")).
		Bold(true)

	optionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#cdd6f4"))

	defaultStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6c7086")).
		Faint(true)

	// show filter if active
	if d.filter != "" {
		b.WriteString(defaultStyle.Render("filter: " + d.filter))
		b.WriteString("\n\n")
	}

	// get visible options
	visibleOptions := d.getVisibleOptions()

	if len(visibleOptions) == 0 {
		b.WriteString(defaultStyle.Render("no matches"))
	} else {
		for _, i := range visibleOptions {
			opt := d.options[i]

			// selection indicator
			prefix := "  "
			if i == d.selected {
				prefix = "> "
			}

			// default indicator
			suffix := ""
			if opt == d.defaultTool {
				suffix = " " + defaultStyle.Render("(default)")
			}

			// render option
			if i == d.selected {
				b.WriteString(prefix + selectedStyle.Render(opt) + suffix)
			} else {
				b.WriteString(prefix + optionStyle.Render(opt) + suffix)
			}
			b.WriteString("\n")
		}
	}

	return openStyle.Render(b.String())
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
