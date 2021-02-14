package algorithms

// time:O(1),space:O(1)
//执行用时: 0 ms 内存消耗: 2.5 MB
func searchNums(nums []int, target int) int {
	l, r, idx := 0, len(nums)-1, -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] <= nums[len(nums)-1] {
			idx = mid
			r = r - 1
		} else {
			l = mid + 1
		}
	}

	res := searchNumsOperate(nums, 0, idx-1, target)
	if res == -1 {
		res = searchNumsOperate(nums, idx, len(nums)-1, target)
	}
	return res
}

func searchNumsOperate(nums []int, l, r, target int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
