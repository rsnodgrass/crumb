package tui

import (
	"strings"
	"testing"
)

func TestRenderHelpOverlay(t *testing.T) {
	// test with standard terminal size
	width := 80
	height := 24

	result := RenderHelpOverlay(width, height)

	// verify output is not empty
	if result == "" {
		t.Error("RenderHelpOverlay returned empty string")
	}

	// verify it contains expected content
	expectedContent := []string{
		"Keyboard Shortcuts",
		"Tab",
		"Next field",
		"Shift+Tab",
		"Previous field",
		"Ctrl+S",
		"Save and exit",
		"Esc",
		"Cancel and exit",
		"/",
		"Open tool selector",
		"?",
		"Toggle this help",
		"Press any key to close",
	}

	for _, content := range expectedContent {
		if !strings.Contains(result, content) {
			t.Errorf("RenderHelpOverlay missing expected content: %q", content)
		}
	}
}

func TestRenderHelpOverlayDifferentSizes(t *testing.T) {
	tests := []struct {
		name   string
		width  int
		height int
	}{
		{"small", 40, 12},
		{"standard", 80, 24},
		{"large", 120, 40},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RenderHelpOverlay(tt.width, tt.height)
			if result == "" {
				t.Errorf("RenderHelpOverlay(%d, %d) returned empty string", tt.width, tt.height)
			}
		})
	}
}
