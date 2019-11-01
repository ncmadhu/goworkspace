package longsubstrworepchr

// LengthOfLongestSubstring - calculates the length of longest substring without repeating characters
func LengthOfLongestSubstring(s string) int {
	length := len(s)
	found := make(map[byte]int, length)
	substrLength := 0
	for start, end := 0, 0; end < length; end++ {
		if val, ok := found[s[end]]; ok {
			start = max(val, start)
		}
		substrLength = max(substrLength, end-start+1)
		found[s[end]] = end + 1
	}
	return substrLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
