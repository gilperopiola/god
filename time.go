package god

import "time"

// Time is a godly concept, as well as his favorite drink.

/* -~-~-~-~ Time -~-~-~-~ */

// Returns a date-string formatted in RFC3339.
// It has date + timezone so that's pretty neat. -> 2006-01-02T15:04:05Z07:00
func FmtDate(date time.Time) string {
	return date.Format(time.RFC3339)
}

// Returns a date-string on a custom format: a substring of time.RubyDate
//
// 		RubyDate 	▶  "Mon Jan 02 15:04:05 -0700 2006"
// 		GodsDate 	▶  "Mon Jan 02 15:04:05"
//
func FmtDateReadable(date time.Time) string {
	return date.Format(time.DateTime)
}

// How many months have passed between 2 dates.
// Does not count unfinished months:
//
// 		date1 = May 10	▶	date2 = April 11	  ▶  0 Months (1 day left)
// 		date1 = May 10	▶	date2 = April 10	  ▶  1 Month
//
func GetMonthsBetween(date1, date2 time.Time) int {

	var (
		yearsSince  = date2.Year() - date1.Year()
		monthsSince = int(date2.Month() - date1.Month())
	)

	if date2.Day() < date1.Day() {
		monthsSince--
	}

	return yearsSince*12 + monthsSince
}

// How many months have passed since a given date.
// Does not count unfinished months:
//
// 		Today = May 10	▶	date = April 11	  ▶  0 Months (1 day left)
// 		Today = May 10	▶	date = April 10	  ▶  1 Month
//
func GetMonthsSince(date time.Time) int {
	return GetMonthsBetween(date, time.Now())
}

// How many days are there in between two dates.
// TODO - Test this to see what's the output when
// before and after are almost the same and when
// they are 23:59:59 apart.
func GetDaysBetween(before, after time.Time) int {
	if before.After(after) {
		before, after = after, before
	}
	return int(after.Sub(before).Hours()) / 24
}
