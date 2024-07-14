package god

import (
	"errors"
	"strings"
)

func GetSubstringBetween(this, that string, in string) (string, error) {

	// Find the start of the first delimiter
	startFrom := strings.Index(in, this)
	if startFrom == -1 {
		return "", errors.New("first delimiter not found")
	}

	// Adjust start index to get content after first delimiter
	startFrom += len(this)

	// Find the start of the second delimiter
	endHere := strings.Index(in, that)
	if endHere == -1 {
		return "", errors.New("second delimiter not found")
	}

	// Extract the content between the delimiters
	if endHere <= startFrom {
		return "", errors.New("second delimiter found before first delimiter")
	}

	return in[startFrom:endHere], nil
}
