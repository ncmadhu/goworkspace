package groupanagram

//GroupAnagrams groups anagram in the given list of strings
func GroupAnagrams(strs []string) [][]string {
	anagramGroups := make(map[[26]byte]int, len(strs))
	result := make([][]string, 0)

	for _, str := range strs {
		list := [26]byte{}
		for _, v := range str {
			list[v-'a']++
		}
		if idx, ok := anagramGroups[list]; ok {
			result[idx] = append(result[idx], str)
		} else {
			result = append(result, []string{str})
			anagramGroups[list] = len(result) - 1
		}
	}
	return result
}
