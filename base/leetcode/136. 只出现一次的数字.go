package leetcode

func singleNumber(nums []int) int {
	// numMap := make(map[int]int)
	// for _, num := range nums {
	// 	numMap[num]++
	// }
	//
	// for k, v := range numMap {
	// 	if v == 1 {
	// 		return k
	// 	}
	// }
	// return 0

	// 解法2: a ^ a = 0, a ^ 0 = a
	single := 0
	for _, num := range nums {
		single ^= num
	}

	return single
}
