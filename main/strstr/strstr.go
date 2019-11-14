package strstr

//StrStr - finds needle in a haystack
func StrStr(haystack string, needle string) int {

	haystackLength := len(haystack)
	needleLength := len(needle)

	if needleLength == 0 {
		return 0
	}

	if haystackLength == 0 || needleLength > haystackLength {
		return -1
	}

	start := -1
	count := 0

	for i, j := 0, 0; i < haystackLength; i++ {

		if j == 0 && haystack[i] == needle[j] {
			start = i
			j++
			count++
			continue
		}
		if j < needleLength && haystack[i] == needle[j] {
			j++
			count++
			continue
		}
		if j == needleLength {
			break
		}
		start = -1
		i = i - j
		j = 0
		count = 0
	}
	if count == needleLength {
		return start
	}
	return -1
}
