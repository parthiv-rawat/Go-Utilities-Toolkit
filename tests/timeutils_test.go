package tests

import (
	"testing"
	"time"
	"go_utilities_toolkit/utils/timeutils"
)

func TestFormatDate(t *testing.T) {
	formatted, err := timeutils.FormatDate("2024-04-01")
	if err != nil || formatted != "01-Apr-2024" {
		t.Errorf("Unexpected date formatting result: %s, err: %v", formatted, err)
	}
}

func TestTimeAgo(t *testing.T) {
	now := time.Now()
	twoHoursAgo := now.Add(-2 * time.Hour)
	ago := timeutils.TimeAgo(twoHoursAgo)
	if ago != "2 hours ago" && ago != "1 hours ago" {
		t.Errorf("Unexpected time ago result: %s", ago)
	}
}