package god

import "testing"

// God writes all his tests in a single file, even if go says otherwise.
// God is above standards.

/* ~-~-~-~-~ Tests: god.go -~-~-~-~ */

/* ~-~-~-~-~ Tests: pope.go -~-~-~-~ */

type ChainableMapAdder[KT comparable, VT any] interface {
	Map(to map[KT]VT, key KT, val VT) ChainableMapAdder[KT, VT]
}

type chainableMap map[int]string

func (cmap chainableMap) Map(i int, s string) chainableMap {
	cmap[i] = s
	return cmap
}

func AddManyToMap[KT comparable, VT any](m map[KT]VT, keysAndVals ...any) map[KT]VT {
	for i := 0; i < len(keysAndVals); i += 2 {
		m[keysAndVals[i].(KT)] = keysAndVals[i+1].(VT)
	}
	return m
}

func TestMapNumToChar(t *testing.T) {

	want := map[int]string{}
	want[1] = "a"
	want[2] = "b"
	want[26] = "z"
	want[-31] = "A"
	want[-6] = "Z"

	want[25] = "y"
	want[-30] = "B"
	want[-7] = "Y"
	want[-1] = "_"
	want[-32] = "@"
	want[0] = "`"
	want[27] = "{"

	for in, want := range want {
		if got := MapNumToChar(in); got != want {
			t.Errorf("🚫 MapNumToChar(%d) returned %s (wanted %s)", in, got, want)
		}
	}

}

/* ~-~-~-~-~ Tests: time.go -~-~-~-~ */
