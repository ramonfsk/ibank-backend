package utils

import "time"

const (
	DATEFORMAT = "2006-01-02"
)

func StringToDate(dateString string) (time.Time, error) {
	return time.Parse(DATEFORMAT, dateString)
}
