package watercontainer

//MaxArea calculates max area from the given set of heights
func MaxArea(height []int) int {
	length := len(height)
	maxArea := 0
	i := 0
	j := length - 1

	for i <= j {
		area := (j - i) * min(height[i], height[j])
		maxArea = max(area, maxArea)
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
