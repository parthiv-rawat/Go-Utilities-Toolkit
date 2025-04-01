package markdown

import (
	"github.com/gomarkdown/markdown"
)

func MarkdownToHTML(input string) string {
	return string(markdown.ToHTML([]byte(input), nil, nil))
}
