package leetcode

func singleNumberII(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	for k, v := range numMap {
		if v == 1 {
			return k
		}
	}
	return 0
}
