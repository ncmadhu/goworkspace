package stringreverse

//Reverse given string
func Reverse(str string) string {
	r := []rune(str)
	for head, tail := 0, len(str)-1; head < len(str)/2; head, tail = head+1, tail-1 {
		r[head], r[tail] = r[tail], r[head]
	}
	return string(r)
}
