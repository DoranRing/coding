package algorithms

// 执行用时: 4 ms 内存消耗: 3.9 MB
func searchRange(nums []int, target int) []int {
	lIdx, rIdx := -1, -1
	// 确定左边界
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid - 1
			if nums[mid] == target {
				lIdx = mid
			}
		} else {
			l = l + 1
		}
	}

	// 确定右边界
	l, r = 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = l + 1
			if nums[mid] == target {
				rIdx = mid
			}
		} else {
			r = mid - 1
		}
	}

	return []int{lIdx, rIdx}
}
