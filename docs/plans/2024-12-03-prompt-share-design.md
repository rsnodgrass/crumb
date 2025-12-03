# prompt-share Design

A beautiful TUI tool for teams to capture and share prompts across AI tools, creating a collaborative learning resource in the repo.

## Overview

### Purpose

Help peer developers learn from each other's AI prompting practices by capturing prompts as-they-go and storing them in a shared `learning/prompts/` directory.

### Core Philosophy

- **Capture-as-you-go** with minimal friction
- **Default everything possible** - user just pastes
- **Beautiful UX** that's enjoyable to use (Catppuccin-inspired)
- **Learning is visible** and shared via git

### Key Decisions

| Decision | Choice |
|----------|--------|
| Tool name | `prompt-share` |
| Built with | Go + Bubble Tea |
| Storage | `learning/prompts/` directory, one markdown file per prompt |
| Config | Global XDG (`~/.config/prompt-share/config.yaml`) |
| Workflow | Single screen, paste-focused |

## Supported AI Tools

Built-in tools (sticky session default, easily overridable):

- Claude Code
- Cursor
- Kiro
- ChatGPT
- Copilot
- Warp AI
- Windsurf
- Aider
- Gemini
- Perplexity

Custom tools can be added via config.

## TUI Design

### Layout

```
┌─────────────────────────────────────────────────────────────┐
│  prompt-share                              [Claude Code ▼]  │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  Prompt:                                                    │
│  ┌─────────────────────────────────────────────────────────┐│
│  │ (paste your prompt here - auto-expands)                 ││
│  │                                                         ││
│  │                                                         ││
│  └─────────────────────────────────────────────────────────┘│
│                                                             │
│  Title: [auto-generated from first line, editable]         │
│                                                             │
│  Tags:  [design] [+]                                        │
│                                                             │
│  Output (optional):                                         │
│  ┌─────────────────────────────────────────────────────────┐│
│  │ (paste AI response if notable)                          ││
│  └─────────────────────────────────────────────────────────┘│
│                                                             │
├─────────────────────────────────────────────────────────────┤
│  [Tab] next field  [Ctrl+S] save  [Esc] cancel  [?] help   │
└─────────────────────────────────────────────────────────────┘
```

### Flow

1. Launch `prompt-share` - cursor in Prompt field
2. Paste prompt (Cmd+V / Ctrl+V)
3. Title auto-generates from first ~60 chars
4. Tool shows session default (top-right dropdown, Tab or click to change)
5. Tags optional - fuzzy autocomplete from existing tags
6. Output field optional - for notable responses
7. Ctrl+S saves - file created - success toast - exit

### Auto-Defaults

| Field | Source |
|-------|--------|
| Timestamp | Current time (ISO 8601) |
| Author | `git config user.name` |
| Tool | Session config default |
| Title | First line of prompt (truncated to ~60 chars) |
| Filename | `YYYY-MM-DD-HH-MM-<slugified-title>.md` |

### Keyboard Shortcuts

| Key | Action |
|-----|--------|
| Tab | Next field |
| Shift+Tab | Previous field |
| Ctrl+S | Save and exit |
| Esc | Cancel and exit |
| `/` or Ctrl+T | Open tool selector |
| ? | Show help |

## File Format

### Filename

```
learning/prompts/2024-12-03-14-32-design-system-color-palette.md
```

### Structure

```markdown
---
title: Design system color palette
tool: ChatGPT
author: Ryan
timestamp: 2024-12-03T14:32:45-08:00
tags:
  - design
  - ux
---

## Prompt

We need to support light and dark modes out of the box. Generate both a light and dark theme.
We should document in our design system that we are heavily inspired by Edward Tufte's data visualization.

## Output

*(optional - included when user pastes a response)*

The recommended palette uses:
- Manushi Blue (#2B82FF) for primary actions...
```

### Frontmatter Fields

| Field | Source | Required |
|-------|--------|----------|
| `title` | Auto from prompt, editable | Yes |
| `tool` | Session default or selected | Yes |
| `author` | `git config user.name` | Yes |
| `timestamp` | Auto (ISO 8601) | Yes |
| `tags` | User-added | No |

## Configuration

### Location

`~/.config/prompt-share/config.yaml` (XDG standard)

### Format

```yaml
# Default tool for new prompts
default_tool: Claude Code

# Custom tools (added to built-in list)
custom_tools:
  - Internal GPT
  - Company Copilot

# Preferred tags (shown first in autocomplete)
favorite_tags:
  - debugging
  - design
  - refactoring
```

## CLI Commands

| Command | Description |
|---------|-------------|
| `prompt-share` | Launch TUI to capture a new prompt |
| `prompt-share readme` | Generate/update `learning/prompts/README.md` |
| `prompt-share config` | Open config file in `$EDITOR` |
| `prompt-share init` | Create `learning/prompts/` directory with starter README |

### Flags

| Flag | Description |
|------|-------------|
| `--tool`, `-t` | Override default tool for this session |
| `--stay` | Don't exit after save (capture multiple prompts) |
| `--version`, `-v` | Show version |
| `--help`, `-h` | Show help |

### Examples

```bash
# Quick capture with default tool
prompt-share

# Capture multiple prompts in a row
prompt-share --stay

# Override tool for this session
prompt-share -t Cursor

# Initialize in a new repo
prompt-share init

# Regenerate README
prompt-share readme
```

## README Generation

Command: `prompt-share readme`

Generates `learning/prompts/README.md`:

```markdown
# Prompt Sharing

Prompts shared by the team to learn from each other.

## Prompts

| Date | Author | Tool | Tags | Title |
|------|--------|------|------|-------|
| 2024-12-03 | Ryan | ChatGPT | design, ux | [Design system color palette](./2024-12-03-14-32-design-system-color-palette.md) |
| 2024-12-03 | Ryan | Claude Code | architecture | [Implement hub management with subagents](./2024-12-03-09-15-implement-hub-management.md) |
| 2024-12-02 | Alex | Cursor | debugging | [Debug auth middleware](./2024-12-02-16-45-debug-auth-middleware.md) |
```

Also outputs a copyable prompt for AI agents to regenerate/customize.

## Visual Design

### Color Palette (Catppuccin-inspired)

| Element | Color | Hex |
|---------|-------|-----|
| Header/title | Lavender | `#b4befe` |
| Labels | Overlay | `#6c7086` |
| Borders | Surface | `#313244` |
| Tool selector | Mauve (accent) | `#cba6f7` |
| Tags | Blue | `#89b4fa` |
| Success | Green | `#a6e3a1` |
| Error | Red | `#f38ba8` |
| Prompt text | Text | `#cdd6f4` |
| Cursor/focus | Peach | `#fab387` |

### Design Principles

- Clean, minimal chrome
- High contrast text for readability
- Subtle borders (not heavy boxes)
- Calm color palette - not distracting
- Smooth cursor transitions
- Input fields expand as content grows
- Tags render as inline badges
- Brief success toast on save

## Project Structure

```
prompt-share/
├── cmd/
│   └── prompt-share/
│       └── main.go           # Entry point
├── internal/
│   ├── tui/
│   │   ├── app.go            # Main Bubble Tea app
│   │   ├── styles.go         # Lip Gloss styles (Catppuccin)
│   │   ├── prompt_view.go    # Main capture screen
│   │   └── components/
│   │       ├── input.go      # Text input/area
│   │       ├── dropdown.go   # Tool selector
│   │       └── tags.go       # Tag input with autocomplete
│   ├── config/
│   │   └── config.go         # XDG config loading
│   ├── storage/
│   │   └── markdown.go       # Read/write prompt files
│   └── readme/
│       └── generator.go      # README.md generation
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

### Dependencies

- `github.com/charmbracelet/bubbletea` - TUI framework
- `github.com/charmbracelet/lipgloss` - Styling
- `github.com/charmbracelet/bubbles` - Input components
- `gopkg.in/yaml.v3` - Config parsing
- `github.com/adrg/xdg` - XDG paths

## Future Considerations

- **MCP/Plugin integration**: AI agents could auto-capture notable prompts
- **Search command**: `prompt-share search <query>` to find prompts
- **Import**: Bulk import from existing prompt history files
- **Export**: Generate shareable formats (PDF, HTML)
