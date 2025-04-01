package web

import (
	"strings"

	"golang.org/x/net/html"
)

func ExtractTitleFromHTML(htmlContent string) string {
	doc, _ := html.Parse(strings.NewReader(htmlContent))
	var crawler func(*html.Node)
	var title string
	crawler = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = n.FirstChild.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			crawler(c)
		}
	}
	crawler(doc)
	return title
}
