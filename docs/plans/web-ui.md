# Web UI Concept Document

**Status**: Concept Only - No Implementation Planned
**Date**: 2025-12-03

---

## 1. Purpose

A future web UI to complement the TUI for browsing and discovering prompts in the shared repository.

## 2. Core Value Proposition

- **Visual browsing** via cards or list views
- **Flexible filtering** by tags, tools, authors, dates
- **Quick access** to full prompt details and metadata
- **One-click copying** to clipboard for use in AI tools
- **Shared knowledge base** treating git repo as source of truth

## 3. Core Screens & Flows

### Library View

Primary browsing interface:

- Grid of prompt cards (title, tool, date, tags)
- List view alternative for dense browsing
- Sorting: by date (default), by tool, by author
- Quick filters sidebar: tags, tools, authors

### Prompt Detail View

- Full prompt text with syntax highlighting
- Metadata panel (author, date, tool, tags)
- Optional output section
- Copy-to-clipboard button
- Link to git history

### Search & Filter

- Full-text search across prompts
- Tag-based filtering (multi-select)
- Tool and author filtering
- Date range filtering

### Collections (Future)

- Group prompts by project/use-case
- User-defined collections
- Auto-collections based on tags

## 4. Technical Approach (Concept Only)

- **Read-only** from git repository
- **Static site generation** (Next.js or Astro)
- **Markdown parsing** with YAML frontmatter
- **Local or hosted** deployment options

## 5. Non-Goals for Initial Version

- No in-browser editing (use TUI or editor)
- No permission model beyond git
- No analytics or tracking
- No user accounts (rely on git identity)
- No real-time collaboration

## 6. Relationship to TUI

| Aspect | TUI | Web UI |
|--------|-----|--------|
| Primary use | Capture prompts | Discover prompts |
| Workflow | Create, edit | Search, filter, copy |
| User base | Contributors | All users (read-only) |

## 7. Future Epic Ideas (NOT for Implementation Now)

- **Epic: Web UI MVP** - Basic browsing (3-4 weeks)
- **Epic: Search & Filter** - Discovery features (2-3 weeks)
- **Epic: Collections** - Organization (2-3 weeks)
- **Epic: Static Deployment** - Hosting (1-2 weeks)

---

**Note**: This document is for planning purposes only. Implementation tasks for the web UI should NOT be created at this time.
