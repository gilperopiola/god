package god

import "time"

// FormatDate formats the given date in RFC3339.
func FormatDate(date time.Time) string {
	return date.Format(time.RFC3339)
}

// GetMonthsSince calculates the months passed since the given date.
func GetMonthsSince(since time.Time) int {
	now := time.Now()
	yearsPassed := now.Year() - since.Year()
	monthsPassed := int(now.Month()) - int(since.Month())
	if now.Day() < since.Day() {
		monthsPassed--
	}
	return yearsPassed*12 + monthsPassed
}

// GetDaysBetween calculates the days in between two dates.
func GetDaysBetween(date1, date2 time.Time) int {
	if date1.After(date2) {
		date1, date2 = date2, date1
	}
	return int(date2.Sub(date1).Hours() / 24)
}
