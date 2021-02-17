package algorithms

//time:O(n),space:O(1)
//执行用时: 4 ms 内存消耗: 2.5 MB
func findMin(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[0] <= nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return nums[l%len(nums)]
}
