package algorithms

import "sort"

//time:O(n*longN),space:O(1)
//执行用时: 72 ms 内存消耗: 6.9 MB
func arrayPairSumBySort(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}
