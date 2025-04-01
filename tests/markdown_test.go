package tests

import (
	"testing"
	"go_utilities_toolkit/utils/markdown"
)

func TestMarkdownToHTML(t *testing.T) {
	input := "# Hello"
	output := markdown.MarkdownToHTML(input)
	if output == "" || output == input {
		t.Errorf("Markdown to HTML conversion failed")
	}
}