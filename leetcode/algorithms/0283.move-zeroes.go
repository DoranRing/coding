package algorithms

// 先将非零元素依次左移，之后的元素置为零
// time: O(n) space:O(1)
// 执行用时: 4 ms 内存消耗: 3.7 MB
func moveZeroesByPoint(nums []int) {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[idx] = nums[i]
			idx++
		}
	}

	for ; idx < len(nums); idx++ {
		nums[idx] = 0
	}
}
