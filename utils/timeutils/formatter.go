package timeutils

import (
	"time"
)

func FormatDate(date string) (string, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}
	return t.Format("02-Jan-2006"), nil
}
