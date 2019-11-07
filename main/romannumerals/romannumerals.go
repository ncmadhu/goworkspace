package romannumerals

import (
	"fmt"
	"strings"
)

//IntToRoman - converts integer to roman numerals
func IntToRoman(num int) string {

	remainder := num
	quotient := num
	roman := ""

	for remainder >= 900 {
		quotient = remainder / 1000
		remainder = remainder % 1000
		roman = roman + strings.Repeat("M", quotient)

		quotient = remainder / 900
		remainder = remainder % 900
		roman = roman + strings.Repeat("CM", quotient)
	}

	for remainder >= 400 {
		quotient = remainder / 500
		remainder = remainder % 500
		roman = roman + strings.Repeat("D", quotient)

		quotient = remainder / 400
		remainder = remainder % 400
		roman = roman + strings.Repeat("CD", quotient)
	}

	for remainder >= 90 {
		quotient = remainder / 100
		remainder = remainder % 100
		roman = roman + strings.Repeat("C", quotient)

		quotient = remainder / 90
		remainder = remainder % 90
		roman = roman + strings.Repeat("XC", quotient)
	}

	for remainder >= 40 {
		quotient = remainder / 50
		remainder = remainder % 50
		roman = roman + strings.Repeat("L", quotient)

		quotient = remainder / 40
		remainder = remainder % 40
		roman = roman + strings.Repeat("XL", quotient)
	}

	for remainder >= 9 {
		quotient = remainder / 10
		remainder = remainder % 10
		roman = roman + strings.Repeat("X", quotient)

		quotient = remainder / 9
		remainder = remainder % 9
		roman = roman + strings.Repeat("IX", quotient)
	}

	for remainder >= 4 {
		quotient = remainder / 5
		remainder = remainder % 5
		roman = roman + strings.Repeat("V", quotient)

		quotient = remainder / 4
		remainder = remainder % 4
		roman = roman + strings.Repeat("IV", quotient)
	}

	roman = roman + strings.Repeat("I", remainder)
	fmt.Println(roman)
	return roman

}
