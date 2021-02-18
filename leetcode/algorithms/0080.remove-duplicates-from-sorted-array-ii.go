package algorithms

// time:O(n),space:O(1)
// 执行用时: 4 ms 内存消耗: 3 MB
func removeDuplicatesByTwoPoint(nums []int) int {
	idx := 1
	for i := idx + 1; i < len(nums); i++ {
		if nums[i] == nums[idx] && nums[idx] == nums[idx-1] {
			continue
		}
		idx++
		nums[idx] = nums[i]
	}

	if len(nums) < idx+1 {
		return len(nums)
	}
	return idx + 1
}
