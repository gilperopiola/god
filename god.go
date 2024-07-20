package god

import (
	"context"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	FailedIntConversionReturns = 0
)

// Used to set default values.
func Configure(failedIntConversionReturns int) {
	FailedIntConversionReturns = failedIntConversionReturns
}

/* -~-~-~-~ Error Handling -~-~-~-~ */

// Shorthand.
func IfErrThen(err error, do func()) {
	if err != nil {
		do()
	}
}

// Use this to ignore the first return value of a function and only get the error.
func OmitReturnedVal(_ any, err error) error {
	return err
}

// Use this to ignore the error of a function and only get the value.
func OmitReturnedErr[T any](val T, _ error) T {
	return val
}

/* -~-~-~-~ Context -~-~-~-~ */

type Ctx = context.Context

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

func ToInt(s string) int {
	if n, err := strconv.Atoi(s); err == nil {
		return n
	}
	return FailedIntConversionReturns
}

// If the error is not nil, we ignore the string and return a default value.
// Helps avoiding an extra if when a func returns (string, error) but you need it as (int, error).
func ToIntAndErr(s string, err error) (int, error) {
	if err != nil {
		return FailedIntConversionReturns, err
	}
	return strconv.Atoi(s)
}

// Is this worth it? Just use int(f) directly.
func F64ToInt(f float64) int {
	return int(f)
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
