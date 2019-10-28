package twosum

//TwoSum to calculate two sum
func TwoSum(nums []int, target int) []int {
	numbers := make(map[int]int, len(nums))
	for i, num := range nums {
		if val, ok := numbers[target-num]; ok {
			return []int{val, i}
		}
		numbers[num] = i
	}
	return nil
}
