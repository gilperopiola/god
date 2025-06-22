package god

import (
	"testing"
	"time"
)

// God writes all his tests in a single file, even if go says otherwise.
// God is above standards.

/* ~-~-~-~-~ Tests: god.go -~-~-~-~ */

/* ~-~-~-~-~ Tests: pope.go -~-~-~-~ */

func TestMapNumToChar(t *testing.T) {
	expected := map[int]string{
		0: returns("`"),
		1: returns("a"), -1: returns("_"),
		2: returns("b"), -2: returns("^"),
		25: returns("y"), -25: returns("G"),
		26: returns("z"), -26: returns("F"),
		27: returns("{"), -27: returns("E"),
	}

	for in, want := range expected {
		if got := MapNumToChar(in); got != want {
			t.Errorf("🚫 MapNumToChar(%d) returned %s (wanted %s)", in, got, want)
		}
	}
}

/* ~-~-~-~-~ Tests: time.go -~-~-~-~ */

func TestGetMonthsBetween(t *testing.T) {
	// Test cases: (date1, date2, expected months between)
	testCases := []struct {
		date1    string
		date2    string
		expected int
	}{
		{"2025-01-01", "2025-06-01", 5},
		{"2025-01-01", "2025-05-31", 4},
		{"2025-01-01", "2025-05-30", 4},
		{"2025-01-01", "2025-05-02", 4},
	}

	for _, tc := range testCases {
		date1, _ := time.Parse("2006-01-02", tc.date1)
		date2, _ := time.Parse("2006-01-02", tc.date2)
		if got := GetMonthsBetween(date1, date2); got != tc.expected {
			t.Errorf("🚫 GetMonthsBetween(%s, %s) returned %d (wanted %d)", tc.date1, tc.date2, got, tc.expected)
		}
	}
}

/* ~-~-~-~-~ Testing Framework -~-~-~-~ */

// Helps reading the test cases
func returns(s string) string {
	return s
}
