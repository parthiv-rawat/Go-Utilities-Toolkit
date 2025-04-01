package tests

import (
	"testing"
	"go_utilities_toolkit/utils/web"
)

func TestExtractTitleFromHTML(t *testing.T) {
	html := "<html><head><title>Test</title></head></html>"
	title := web.ExtractTitleFromHTML(html)
	if title != "Test" {
		t.Errorf("Expected 'Test', got '%s'", title)
	}
}