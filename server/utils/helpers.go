package utils

import "time"

func StringToDate(dateString string) (time.Time, error) {
	return time.Parse(DATEFORMAT, dateString)
}

func StatusAsText(status int) string {
	if status == 0 {
		return INACTIVE
	}

	return ACTIVE
}
