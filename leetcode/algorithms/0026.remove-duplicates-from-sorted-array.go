package algorithms

// 双指针
// time: O(n) space: O(1)
// 执行用时: 0 ms 内存消耗: 4.6 MB
func removeDuplicatesByPoint(nums []int) int {
	idx := 1
	for i := idx; i < len(nums); i++ {
		if nums[idx-1] < nums[i] {
			nums[idx] = nums[i]
			idx++
		}
	}
	if idx > len(nums) {
		return len(nums)
	}
	return idx
}
