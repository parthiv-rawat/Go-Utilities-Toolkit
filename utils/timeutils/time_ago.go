package timeutils

import (
	"fmt"
	"time"
)

func TimeAgo(t time.Time) string {
	duration := time.Since(t)
	switch {
	case duration < time.Minute:
		return "Just now"
	case duration < time.Hour:
		return fmt.Sprintf("%d minutes ago", int(duration.Minutes()))
	case duration < 24*time.Hour:
		return fmt.Sprintf("%d hours ago", int(duration.Hours()))
	default:
		return t.Format("02 Jan 2006")
	}
}
