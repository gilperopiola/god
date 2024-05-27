package god

import (
	"os"
	"strings"
)

func CreateFile(filename string, contents string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	if _, err = file.WriteString(contents); err != nil {
		return err
	}

	return nil
}

func GetFileExtension(filename string) string {
	split := strings.Split(filename, ".")
	return "." + split[len(split)-1]
}
