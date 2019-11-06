package atoi

import (
	"fmt"
	"strconv"
	"unicode"
)

// Atoi - converts string to integer
func Atoi(str string) int32 {
	length := len(str)
	runes := []rune(str)
	start, end := 0, 0
	foundStart := false
	for i, c := range runes {
		if unicode.IsSpace(c) {
			continue
		} else if unicode.IsDigit(c) {
			start = i
			foundStart = true
			break
		} else if c == '-' || c == '+' && i+1 < length && unicode.IsDigit(runes[i+1]) {
			start = i
			foundStart = true
			break
		} else {
			break
		}
	}
	if foundStart {
		end = start
		for i := start + 1; i < length; i++ {
			if !unicode.IsDigit(runes[i]) {
				break
			}
			end++
		}
		number, _ := strconv.ParseInt(str[start:end+1], 10, 32)
		fmt.Println(number)
		return int32(number)
	}
	return 0
}
