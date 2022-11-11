package common

import "time"

func ConvertStringtoDate(dateString string) (time.Time, error) {
	date, error := time.Parse("2006-01-02", dateString)
	return date, error
}
func ConvertStringtoTime(dateString string) (time.Time, error) {
	date, error := time.Parse("10:20", dateString)
	return date, error
}
