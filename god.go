package god

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

/* -~-~-~-~ God -~-~-~-~ */

// 👀

/* -~-~-~-~ Types -~-~-~-~ */

type Ctx = context.Context

type Num interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

/* -~-~-~-~ Error Handling -~-~-~-~ */

// Shorthand.
func IfErrThen(err error, do func()) {
	if err != nil {
		do()
	}
}

// Use this to ignore the first return value of a function and only get the error.
func OnlyGetErr(_ any, err error) error {
	return err
}

// Use this to ignore the error of a function and only get the value.
func DontGetErr[T any](val T, _ error) T {
	return val
}

/* -~-~-~-~ Context -~-~-~-~ */

func NewCtx() Ctx {
	return context.Background()
}

func NewCtxWithTimeout(d time.Duration) (Ctx, context.CancelFunc) {
	return context.WithTimeout(NewCtx(), d)
}

/* -~-~-~-~ Type Conversions -~-~-~-~ */

func ToString(n int) string {
	return strconv.Itoa(n)
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

/* -~-~-~-~ Strings -~-~-~-~ */

// Gets the substring between two delimiters inside of a given string.
//
//	in [+10] 		▶  this [       ▶  that ]       ▶  "+10"
//	in <p>Asd</p> 	▶  this <p>     ▶  that </p>    ▶  "Asd"
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

/* -~-~-~-~ Environment Variables -~-~-~-~ */

// Use this to get env vars. You can provide a fallback value in case it doesn't exist.
func GetEnv[T string | bool | int](key string, or T) T {
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

/* -~-~-~-~ Pary 🙏 -~-~-~-~ */

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
