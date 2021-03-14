package algorithms

import "sort"

//time:O(n*logN), space:O(1)
//执行用时: 56 ms 内存消耗: 6.6 MB
func maximumProductBySort(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	res1 := nums[length-1] * nums[length-2] * nums[length-3]
	res2 := nums[0] * nums[1] * nums[length-1]
	if res1 > res2 {
		return res1
	}
	return res2
}
