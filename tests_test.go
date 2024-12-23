package god

import "testing"

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

/* ~-~-~-~-~ Testing Framework -~-~-~-~ */

// Helps reading the test cases
func returns(s string) string {
	return s
}
