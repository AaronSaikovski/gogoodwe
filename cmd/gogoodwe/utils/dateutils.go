package utils

import "time"

func isWithin7Days(dateStr1, dateStr2 string) (bool, error) {
	layout := "2006-01-02 15:04"

	date1, err := time.Parse(layout, dateStr1)
	if err != nil {
		return false, err
	}

	date2, err := time.Parse(layout, dateStr2)
	if err != nil {
		return false, err
	}

	diff := date2.Sub(date1).Abs()
	return diff <= 7*24*time.Hour, nil
}
