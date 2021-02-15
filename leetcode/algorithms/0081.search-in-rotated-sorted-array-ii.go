package algorithms

// time:O(n),space:O(1)
//执行用时: 4 ms 内存消耗: 3.2 MB
func searchInRotatedSortedArrayIIByLoop(nums []int, target int) bool {
	for _, num := range nums {
		if target == num {
			return true
		}
	}
	return false
}

// time:O()
// 执行用时: 8 ms 内存消耗: 3.2 MB
func searchInRotatedSortedArrayIIByBinarySearch(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}

		// 确定不了哪边是有序，排除掉左边边界
		if nums[l] == nums[mid] {
			l++
		} else if nums[l] < nums[mid] {
			// 左边排序部分
			if nums[l] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			// 右边排序部分
			if nums[r] >= target && target > nums[mid] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
