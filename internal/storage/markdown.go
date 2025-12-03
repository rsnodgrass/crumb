package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type PromptMetadata struct {
	Title       string
	Description string
	Tags        []string
	Author      string
	CreatedAt   time.Time
}

type MarkdownStorage struct {
	baseDir string
}

func NewMarkdownStorage(baseDir string) *MarkdownStorage {
	return &MarkdownStorage{
		baseDir: baseDir,
	}
}

func (m *MarkdownStorage) Save(metadata PromptMetadata, content string) error {
	if err := os.MkdirAll(m.baseDir, 0755); err != nil {
		return fmt.Errorf("failed to create prompts directory: %w", err)
	}

	filename := m.generateFilename(metadata.Title)
	filepath := filepath.Join(m.baseDir, filename)

	markdown := m.formatMarkdown(metadata, content)

	if err := os.WriteFile(filepath, []byte(markdown), 0644); err != nil {
		return fmt.Errorf("failed to write prompt file: %w", err)
	}

	return nil
}

func (m *MarkdownStorage) generateFilename(title string) string {
	// Convert title to kebab-case filename
	safe := strings.ToLower(title)
	safe = strings.ReplaceAll(safe, " ", "-")
	safe = strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' {
			return r
		}
		return -1
	}, safe)
	return safe + ".md"
}

func (m *MarkdownStorage) formatMarkdown(metadata PromptMetadata, content string) string {
	var sb strings.Builder

	sb.WriteString("# ")
	sb.WriteString(metadata.Title)
	sb.WriteString("\n\n")

	if metadata.Description != "" {
		sb.WriteString(metadata.Description)
		sb.WriteString("\n\n")
	}

	sb.WriteString("**Tags:** ")
	if len(metadata.Tags) > 0 {
		sb.WriteString(strings.Join(metadata.Tags, ", "))
	} else {
		sb.WriteString("(none)")
	}
	sb.WriteString("\n\n")

	sb.WriteString("**Author:** ")
	sb.WriteString(metadata.Author)
	sb.WriteString("\n\n")

	sb.WriteString("**Created:** ")
	sb.WriteString(metadata.CreatedAt.Format("2006-01-02 15:04:05"))
	sb.WriteString("\n\n")

	sb.WriteString("---\n\n")
	sb.WriteString(content)
	sb.WriteString("\n")

	return sb.String()
}
