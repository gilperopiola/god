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
