package algorithms

// 执行用时: 4 ms 内存消耗: 3.9 MB
func searchRange(nums []int, target int) []int {
	lIdx := BinarySearchFirst(nums, target)
	rIdx := BinarySearchLast(nums, target)
	return []int{lIdx, rIdx}
}

func BinarySearchFirst(nums []int, target int) int {
	idx := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			r = mid - 1
			if nums[mid] == target {
				idx = mid
			}
		} else {
			l = l + 1
		}
	}
	return idx
}

func BinarySearchLast(nums []int, target int) int {
	idx := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= target {
			l = l + 1
			if nums[mid] == target {
				idx = mid
			}
		} else {
			r = mid - 1
		}
	}
	return idx
}
