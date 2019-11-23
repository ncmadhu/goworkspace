package minimumwindowsubstring

import "fmt"

//MinWindow - finds the minimum window substring of given string
func MinWindow(s string, t string) string {

	var result string
	have := [128]int{}
	need := [128]int{}
	srcLen := len(s)
	targetLen := len(t)
	currentLength := srcLen + 1

	for _, char := range t {
		fmt.Println(char)
		need[char]++
	}

	for left, right, found := 0, 0, 0; right < srcLen; right++ {
		if have[s[right]] < need[s[right]] {
			found++
		}
		have[s[right]]++

		for left <= right && have[s[left]] > need[s[left]] {
			have[s[left]]--
			left++
		}

		length := right - left + 1
		if found == targetLen && currentLength > length {
			currentLength = length
			result = s[left : right+1]
		}
	}
	return result

}
