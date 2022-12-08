package loop

func kLoop(nums []int, step int) {
	l := len(nums)
	for i := 0; i < step; i++ {
		for j := i; j < l; j += step {
			nums[j] = 4
		}
	}
}

func CreateSource(len int) []int {
	nums := make([]int, 0, len)

	for i := 0; i < len; i++ {
		nums = append(nums, i)
	}

	return nums
}
