package number

import (
	"fmt"
	"math"
	"strconv"
)

// AddTimestampFraction add last N digit of timestamp as value's decimal number
func AddTimestampFraction(value int, unixTimestamp int64, digit int) float64 {
	// validate the digit value is not breaking the logic
	if digit <= 0 || digit > CountDigits(unixTimestamp) {
		return float64(value)
	}

	// generate modulus number
	digitScale := math.Pow(10, float64(digit))

	// get last N digit for unix timestamp
	lastDigits := int(unixTimestamp) % int(digitScale)

	// change into decimal number
	fraction := float64(lastDigits) / float64(digitScale)

	return float64(value) + fraction
}

// CountDigits used to get the number of digit
func CountDigits(n int64) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n // convert negative numbers
	}
	return int(math.Log10(float64(n))) + 1
}

// RoundUp is function for round up from division value
func RoundUp(number, dividedBy int) int {
	divide := number / dividedBy
	restFor := number % dividedBy
	result := divide
	if restFor > 0 {
		result++
	}

	return result
}

// OrdinalSuffix to get ordinal suffix from given number
func OrdinalSuffix(n int) string {
	if n <= 0 {
		return fmt.Sprintf("%d", n)
	}

	// Special case for 11, 12, 13
	if n%100 >= 11 && n%100 <= 13 {
		return fmt.Sprintf("%dth", n)
	}

	switch n % 10 {
	case 1:
		return fmt.Sprintf("%dst", n)
	case 2:
		return fmt.Sprintf("%dnd", n)
	case 3:
		return fmt.Sprintf("%drd", n)
	default:
		return fmt.Sprintf("%dth", n)
	}
}

// GenerateIndoNumFormat to generate indonesian number format using . as demiliter. ex : 1.000.000
func GenerateIndoNumFormat(num int64) string {
	in := strconv.FormatInt(num, 10)
	numOfDigits := len(in)
	if num < 0 {
		numOfDigits-- // First character is the - sign (not a digit)
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if num < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = '.'
		}
	}
}

// RoundFloat to round a float to flexible decimal places
func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
