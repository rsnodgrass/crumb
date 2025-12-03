package storage

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

// Save writes markdown content to a file in the base directory.
// Returns the full filepath on success or an error.
func (m *MarkdownStorage) Save(filename string, content string) (string, error) {
	if err := os.MkdirAll(m.baseDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	fullPath := filepath.Join(m.baseDir, filename)
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	return fullPath, nil
}

// SaveWithMetadata is the legacy method for saving with structured metadata
func (m *MarkdownStorage) SaveWithMetadata(metadata PromptMetadata, content string) error {
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

// GetFrequentTags scans all crumb files and returns tags sorted by frequency
func (m *MarkdownStorage) GetFrequentTags(limit int) []string {
	tagCounts := make(map[string]int)

	entries, err := os.ReadDir(m.baseDir)
	if err != nil {
		return []string{}
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") || entry.Name() == "README.md" {
			continue
		}

		tags := m.extractTagsFromFile(filepath.Join(m.baseDir, entry.Name()))
		for _, tag := range tags {
			tagCounts[tag]++
		}
	}

	// sort by frequency
	type tagFreq struct {
		tag   string
		count int
	}
	freqs := make([]tagFreq, 0, len(tagCounts))
	for tag, count := range tagCounts {
		freqs = append(freqs, tagFreq{tag, count})
	}
	sort.Slice(freqs, func(i, j int) bool {
		return freqs[i].count > freqs[j].count
	})

	// return top N
	result := make([]string, 0, limit)
	for i := 0; i < len(freqs) && i < limit; i++ {
		result = append(result, freqs[i].tag)
	}
	return result
}

// extractTagsFromFile parses YAML frontmatter to extract tags
func (m *MarkdownStorage) extractTagsFromFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	inFrontmatter := false
	inTags := false
	tags := []string{}

	for scanner.Scan() {
		line := scanner.Text()

		// detect frontmatter boundaries
		if line == "---" {
			if !inFrontmatter {
				inFrontmatter = true
				continue
			} else {
				break // end of frontmatter
			}
		}

		if !inFrontmatter {
			continue
		}

		// detect tags section
		if strings.HasPrefix(line, "tags:") {
			inTags = true
			// check for inline tags: tags: [tag1, tag2]
			rest := strings.TrimPrefix(line, "tags:")
			rest = strings.TrimSpace(rest)
			if strings.HasPrefix(rest, "[") && strings.HasSuffix(rest, "]") {
				// inline array format
				inner := strings.Trim(rest, "[]")
				for _, t := range strings.Split(inner, ",") {
					t = strings.TrimSpace(t)
					if t != "" {
						tags = append(tags, t)
					}
				}
				inTags = false
			}
			continue
		}

		// parse tag list items
		if inTags {
			if strings.HasPrefix(line, "  - ") {
				tag := strings.TrimPrefix(line, "  - ")
				tag = strings.TrimSpace(tag)
				if tag != "" {
					tags = append(tags, tag)
				}
			} else if !strings.HasPrefix(line, "  ") && !strings.HasPrefix(line, "\t") {
				// no longer indented, end of tags
				inTags = false
			}
		}
	}

	return tags
}
