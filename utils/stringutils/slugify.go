package stringutils

import (
	"regexp"
	"strings"
)

func Slugify(str string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	slug := strings.ToLower(re.ReplaceAllString(str, "-"))
	return strings.Trim(slug, "-")
}
