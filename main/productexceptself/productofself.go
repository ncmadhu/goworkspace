package productexceptself

//ProductExceptSelf - product except self
func ProductExceptSelf(nums []int) []int {
	length := len(nums)
	output := make([]int, length)
	product := 1
	for i := 0; i < length; i++ {
		output[i] = product
		product = product * nums[i]
	}
	product = 1
	for i := length - 1; i >= 0; i-- {
		output[i] = output[i] * product
		product = product * nums[i]
	}
	return output
}
