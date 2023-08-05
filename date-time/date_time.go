package date_time

import (
	"strconv"
	"time"
)

func UnixToTime(unix int64) time.Time {
	str := strconv.Itoa(int(unix))

	if len(str) == 13 { // unix with millisecond
		return time.UnixMilli(unix)
	}

	return time.Unix(unix, 0)
}

func IsSame(t1, t2 time.Time) bool {
	y1, m1, d1 := t1.Date()
	y2, m2, d2 := t2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func ParseISODate(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, date) // ISO Date
}

func ToISODate(date time.Time) string {
	return date.Format(time.RFC3339)
}
