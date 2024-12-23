package god

import (
	"os"
	"strings"
)

/* ~-~-~-~-~ The Mighty Pope -~-~-~-~ */

// Not as useful as God. But hey, still the Pope.

/* ~-~-~-~-~ Slices & Maps -~-~-~-~ */

// Mainly used when you have a function that accepts an optional param
func GetFirstOrZeroVal[T any](slice ...T) T {
	if len(slice) == 0 {
		return *(new(T))
	}
	return slice[0]
}

func AddToMap[KT comparable, VT any](m map[KT]VT, kv ...any) map[KT]VT {
	for i := 0; i < len(kv); i += 2 {
		m[kv[i].(KT)] = kv[i+1].(VT)
	}
	return m
}

/* ~-~-~-~-~ Encryption -~-~-~-~-~ */

// Returns the letter in the alphabet's position of the given number.
//
//	n = 1  ▶ MapNumToChar(1)  = "a"
//	n = 2  ▶ MapNumToChar(2)  = "b"
//	n = 10 ▶ MapNumToChar(10) = "j"
func MapNumToChar(n int) string {
	return string(rune(asciiLowerA + n))
}

const asciiLowerA = 96

/* ~-~-~-~-~ OS | Filesystem | I/O ~-~-~-~-~ */

func CreateFile(filename, content string) error {
	file, err := os.Create(filename)
	if err != nil || OnlyGetErr(file.WriteString(content)) != nil {
		return err
	}
	return nil
}

func GetFileExt(filename string) string {
	split := strings.Split(filename, ".")
	return "." + split[len(split)-1]
}
