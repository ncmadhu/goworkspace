package threesumclosest

import (
	"math"
	"sort"
)

//ThreeSumClosest - find the numbers added closest to the target
func ThreeSumClosest(nums []int, target int) int {
	length := len(nums)
	sort.Ints(nums)
	diff := math.MaxInt32
	result := -target
	if length < 3 {
		return 0
	}
	for i := 0; i < length; i++ {
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
			if sum == target {
				return target
			} else if sum < target {
				currentDiff := target - sum
				if currentDiff < diff {
					diff = currentDiff
					result = sum
				}
				j++
			} else {
				currentDiff := sum - target
				if currentDiff < diff {
					diff = currentDiff
					result = sum
				}
				k--
			}
		}
	}
	return result
}
