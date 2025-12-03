# TUI UX Design Principles

## Vision

A 2025-grade terminal user interface that developers love - minimal friction, capture-as-you-go, terminal-native.

## Core Principles

### 1. Minimal Friction

**From `prompt-share` to saved prompt in seconds**

- No modal labyrinths or unnecessary confirmations
- Quick flow: Title → Tool → Tags → Prompt → Ctrl+S
- Smart defaults reduce required input
- Exit immediately after save (unless `--stay` flag)

### 2. Single Clear Focus

**Strong visual feedback for current field**

- Peach accent border (#fab387) around focused component
- Arrow indicator (→) before focused field label
- Logical tab order matching natural workflow

### 3. Micro-copy & Inline Help

**Context-aware guidance without clutter**

- One-line description at top explaining purpose
- Context-sensitive key hints in footer
- `?` for full help overlay with all shortcuts

### 4. Visual Calm

**Catppuccin Mocha palette for comfortable extended use**

- Subtle borders and generous whitespace
- Essential information only, no visual clutter
- Muted colors for labels, bright colors for focus

### 5. Fast Workflows

**Optimize for speed without sacrificing clarity**

- Snappy tag autocomplete with favorites first
- `/` to jump directly to tool selector
- Type-to-filter in dropdowns
- Number keys (1-5) for quick tag selection

### 6. Great Defaults

**Reduce cognitive load with smart configuration**

- Pre-select `default_tool` from config
- Surface `favorite_tags` as quick chips
- Sensible `output_dir` (learning/prompts)

## Target User Flow

1. **Launch:** `prompt-share`
2. **(Optional)** Edit title if auto-generated isn't right
3. **(Optional)** Change tool with `/`
4. **(Optional)** Add tags (type or press 1-5)
5. Paste/type prompt
6. **Ctrl+S** to save
7. See success toast with filepath
8. Exit (or continue if `--stay`)

**Time target:** < 30 seconds from launch to saved prompt

## Keyboard Shortcuts

| Key | Action | Context |
|-----|--------|---------|
| Tab | Next field | Always |
| Shift+Tab | Previous field | Always |
| Ctrl+S | Save and exit | Always |
| Esc | Cancel and exit | Always |
| / | Open tool selector | Always |
| ? | Show help overlay | Always |
| Enter | Add tag | Tags field |
| 1-5 | Quick-select suggestion | Tags field |
| ↑/↓ | Navigate list | Dropdown open |

## Color Vocabulary

| Element | Color | Hex |
|---------|-------|-----|
| Primary | Lavender | #b4befe |
| Accent/Focus | Peach | #fab387 |
| Success | Green | #a6e3a1 |
| Error | Red | #f38ba8 |
| Muted | Overlay | #6c7086 |
| Text | Text | #cdd6f4 |
| Border | Surface | #313244 |
| Background | Base | #1e1e2e |

## Layout Structure

```
┌─────────────────────────────────────────────────┐
│ prompt-share                                     │
│ Capture prompts to learning/prompts/            │
├─────────────────────────────────────────────────┤
│ → Title: [Your prompt title here________]       │
│   Tool:  [Claude Code ▼]                        │
│   Tags:  [coding] [debug] [+add]                │
│          Suggestions: 1:testing 2:design        │
│   Prompt:                                       │
│   ┌───────────────────────────────────────┐     │
│   │ Write your prompt here...              │     │
│   └───────────────────────────────────────┘     │
│   Output (optional):                            │
│   ┌───────────────────────────────────────┐     │
│   │ AI response...                         │     │
│   └───────────────────────────────────────┘     │
├─────────────────────────────────────────────────┤
│ Tab: next • Ctrl+S: save • Esc: quit • ?: help  │
└─────────────────────────────────────────────────┘
```

---

**Last Updated:** 2025-12-03
