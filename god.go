package god

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// 👀 ─ God is always watching.

type Num interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

/* ── ── ── Error Handling ── ── ── */

// Shorthand.
func IfErrThen(err error, do func()) {
	if err != nil {
		do()
	}
}

// Use this to ignore the first return value of a function and only get the error.
func SkipVal[T any](_ T, err error) error {
	return err
}

// Use this to ignore the error of a function and only get the value.
func SkipErr[T any](val T, _ error) T {
	return val
}

/* ── ── ── Context ── ── ── */

type Ctx = context.Context

func NewCtx() Ctx {
	return context.Background()
}

func NewCtxWithTimeout(dur time.Duration) (Ctx, context.CancelFunc) {
	return context.WithTimeout(NewCtx(), dur)
}

/* ── ── ── Type Conversions ── ── ── */

func ToString[T any](value T) string {
	switch val := any(value).(type) {
	case string:
		return val
	case bool:
		return strconv.FormatBool(val)
	case int:
		return strconv.Itoa(val)
	case uint, int8, int16, int32, int64, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		return fmt.Sprintf("%.2f", val)
	case time.Time:
		return val.Format(time.RFC3339)
	case []byte:
		return string(val)
	default:
		return fmt.Sprintf("%v", val)
	}
}

// Note: The default value if the conversion to int fails is 0.
// If you need to differentiate between a 0 and a failed conversion
// you're out of luck.
func ToInt(s string, fallback ...int) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return 0
}

// If the error is not nil, we ignore the string and default to 0.
// Helps avoiding an extra <if> when a func returns (string, error) but you need it as (int, error).
func ToIntAndErr(s string, err error) (int, error) {
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(s)
}

// This can overflow with numbers greater than the OS's max int size.
func ToIntSlice[T Num](numbers []T) []int {
	intNumbers := make([]int, len(numbers))
	for i, number := range numbers {
		intNumbers[i] = int(number)
	}
	return intNumbers
}

// Gets the keys of a map.
func GetMapKeys[T any](m map[string]T) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

/* ── ── ── Strings ── ── ── */

// Gets the substring between two delimiters inside of a given string.
//
//	in [+10] 		▶  this [       ▶  that ]       ──▶  "+10"
//	in <p>Asd</p> 	▶  this <p>     ▶  that </p>    ──▶  "Asd"
func SubstrBetween(this, that string, in string) (string, error) {

	startFrom := strings.Index(in, this) // Get start of the first delimiter, then move to the end of it
	if startFrom == -1 {
		return "", errors.New("first delimiter not found")
	}
	startFrom += len(this)

	endHere := strings.Index(in, that) // Get start of the second delimiter
	if endHere == -1 {
		return "", errors.New("second delimiter not found")
	}

	if endHere <= startFrom {
		return "", errors.New("second delimiter found before the first one")
	}

	return in[startFrom:endHere], nil
}

/* ── ── ── Env Vars ── ── ── */

// Use this to get env vars. You can provide a fallback value in case it doesn't exist.
func GetEnv[T int | string | bool](key string, or T) T {
	val, ok := os.LookupEnv(key)
	if !ok {
		return or
	}

	switch any(or).(type) {
	case string:
		return any(val).(T)
	case bool:
		return any(strings.ToLower(val) == "true" || val == "1").(T)
	case int:
		if intVal, err := strconv.Atoi(val); err == nil {
			return any(intVal).(T)
		}
	}

	return or
}

/* ── ── ── Pary 🙏 ── ── ── */

// ))
//  °\°
//   _\
//  |
// 👄         I am God. Pls pary.
// /

var TimesParied = 0

func Pary() {
	TimesParied++
}
