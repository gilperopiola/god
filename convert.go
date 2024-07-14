package god

import (
	"log"
	"strconv"
)

func ToInt(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	log.Println("string to int failed")
	return -1
}

func ToString(n int) string {
	return strconv.Itoa(n)
}

func Float64ToInt(f float64) int {
	return int(f)
}

/* -~-~-~-~-~-~-~-~-~-~-~-~-~-~-~-~ */

// Helps avoiding an extra if when a func returns (string, error) but you need it as a string
func ToIntAndErr(s string, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}
