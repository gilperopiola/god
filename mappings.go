package god

func MapIntToLetter(n int) string {
	return string(rune(n + 95)) // 96 -> ASCII 'a'
}
