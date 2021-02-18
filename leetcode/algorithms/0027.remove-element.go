package algorithms

// time:O(n),space:O(1)
//执行用时: 0 ms 内存消耗: 2.1 MB
func removeElementByTwoPoint(nums []int, val int) int {
	idx := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			idx++
			nums[idx] = nums[i]
		}
	}
	return idx + 1
}
