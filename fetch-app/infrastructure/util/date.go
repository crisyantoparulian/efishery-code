package util

import (
	"time"
)

func ForceParseDate(str string) (time.Time, error) {
	var (
		date time.Time
		err  error
	)
	layouts := []string{time.RFC3339, time.RFC822, "2 Jan 2006 15:04:05", "2006/1/2 15:04:05", "2006-01-02 15:04:05.000",
		"2006/1/2 15:4:5", time.RFC3339Nano, "2006-01-02T15:04:05.000-0700", "2/1/2006 15:04:05", "2006-01-02 15:04:05"}
	layoutsLen := len(layouts) - 1
	for k, l := range layouts {
		date, err = time.Parse(l, str)
		if err != nil && k == layoutsLen {
			return date, err
		} else if err == nil {
			return date, nil
		}

	}
	return date, nil
}
