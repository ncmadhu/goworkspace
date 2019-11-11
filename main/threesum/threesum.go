package threesum

import (
	"sort"
)

//ThreeSum - Find the three numbers which add to the given value
func ThreeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	result := [][]int{}
	if length < 3 {
		return result
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, length-1
		for j < k {
			if j > i+1 && nums[j] == nums[j-1] {
				j++
				continue
			}
			if k < length-1 && nums[k] == nums[k+1] {
				k--
				continue
			}
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--
			} else if sum < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result
}
