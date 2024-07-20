package god

import (
	"os"
	"strings"
)

/* ~-~-~-~-~ The Mighty Pope -~-~-~-~ */

// Not as useful as God. But still the Pope.

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

func CreateFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil || OmitReturnedVal(file.WriteString(content)) != nil {
		return err
	}
	return nil
}

func GetFileExt(filename string) string {
	split := strings.Split(filename, ".")
	return "." + split[len(split)-1]
}
