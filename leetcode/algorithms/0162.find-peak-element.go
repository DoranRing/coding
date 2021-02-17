package algorithms

//
//执行用时: 4 ms 内存消耗: 2.7 MB
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if mid < len(nums)-1 && nums[mid] < nums[mid+1] {
			l++
		} else {
			r = mid
		}
	}

	if l != len(nums) {
		return l
	}
	return -1
}
