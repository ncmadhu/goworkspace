package missingnumber

//MissingNumber - finds the missing number
func MissingNumber(nums []int) int {
	count := 0
	sum := 0
	actualSum := 0
	for _, num := range nums {
		actualSum = actualSum + num
		sum = sum + count
		count++
	}
	sum = sum + count
	MissingNumber := sum - actualSum
	return MissingNumber

}
