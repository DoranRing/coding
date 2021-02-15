package algorithms

// time:O(n),space:O(1)
// 执行用时: 4 ms 内存消耗: 2.5 MB
func searchInRotatedSortedArrayByLoop(nums []int, target int) int {
	for idx, num := range nums {
		if target == num {
			return idx
		}
	}
	return -1
}

// time:O(1),space:O(1)
//执行用时: 4 ms 内存消耗: 2.5 MB
func searchInRotatedSortedArrayByBinarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			// 左边排序部分
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右边排序部分
			if nums[len(nums)-1] >= target && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
