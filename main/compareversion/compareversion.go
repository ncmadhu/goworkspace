package compareversion

import (
	"strconv"
	"strings"
)

//CompareVersion - compares version1 with version2
func CompareVersion(version1 string, version2 string) int {
	ver1Arr := strings.Split(version1, ".")
	ver2Arr := strings.Split(version2, ".")
	sign := 1
	result := 0
	count := 0

	if len(ver2Arr) < len(ver1Arr) {
		ver1Arr, ver2Arr = ver2Arr, ver1Arr
		sign = -1
	}

	for i, element := range ver1Arr {
		count++
		ver1, _ := strconv.Atoi(element)
		ver2, _ := strconv.Atoi(ver2Arr[i])
		if ver1 > ver2 {
			result = 1
			break
		} else if ver1 < ver2 {
			result = -1
			break
		}
	}
	if result == 0 {
		for count < len(ver2Arr) {
			ver, _ := strconv.Atoi(ver2Arr[count])
			if ver > 0 {
				result = -1
				break
			}
			count++
		}
	}
	return result * sign

}
