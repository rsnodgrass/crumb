package readme

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type Prompt struct {
	Title       string
	Description string
	Tags        []string
	Filename    string
}

type Generator struct {
	promptsDir string
}

func NewGenerator(promptsDir string) *Generator {
	return &Generator{
		promptsDir: promptsDir,
	}
}

func (g *Generator) Generate() error {
	prompts, err := g.scanPrompts()
	if err != nil {
		return fmt.Errorf("failed to scan prompts: %w", err)
	}

	readmeContent := g.formatReadme(prompts)
	readmePath := filepath.Join(g.promptsDir, "README.md")

	if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
		return fmt.Errorf("failed to write README: %w", err)
	}

	return nil
}

func (g *Generator) scanPrompts() ([]Prompt, error) {
	entries, err := os.ReadDir(g.promptsDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []Prompt{}, nil
		}
		return nil, err
	}

	prompts := make([]Prompt, 0)
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".md") || entry.Name() == "README.md" {
			continue
		}

		prompt := Prompt{
			Filename: entry.Name(),
			Title:    strings.TrimSuffix(entry.Name(), ".md"),
		}
		prompts = append(prompts, prompt)
	}

	return prompts, nil
}

func (g *Generator) formatReadme(prompts []Prompt) string {
	var sb strings.Builder

	sb.WriteString("# Prompt Library\n\n")
	sb.WriteString("Collection of prompts for various tasks.\n\n")

	if len(prompts) == 0 {
		sb.WriteString("No prompts available yet.\n")
		return sb.String()
	}

	sb.WriteString("## Available Prompts\n\n")

	for _, prompt := range prompts {
		sb.WriteString("- [")
		sb.WriteString(prompt.Title)
		sb.WriteString("](")
		sb.WriteString(prompt.Filename)
		sb.WriteString(")\n")
	}

	return sb.String()
}

// Generate is a convenience function that creates a generator and generates the README
func Generate(promptsDir string) (string, error) {
	g := NewGenerator(promptsDir)
	prompts, err := g.scanPrompts()
	if err != nil {
		return "", fmt.Errorf("failed to scan prompts: %w", err)
	}
	return g.formatReadme(prompts), nil
}
