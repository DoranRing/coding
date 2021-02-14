package algorithms

import "sort"

// 排序，然后使用双指针对应第一、第二个元素，然后通过二分查找确认第三个元素
// time:O(n2), space:O(1)
//执行用时: 420 ms 内存消耗: 7.3 MB
func threeSumBySortAndTwoPoint(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			need := 0 - (nums[i] + nums[j])
			idx := search(nums, j+1, len(nums)-1, need)
			if idx != -1 && !contains(res, nums, i, j, idx) {
				res = append(res, []int{nums[i], nums[j], nums[idx]})
			}
		}
	}

	return res
}

func search(nums []int, l, r int, val int) int {
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == val {
			return mid
		} else if nums[mid] < val {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func contains(res [][]int, nums []int, i, j, k int) bool {
	for _, item := range res {
		if item[0] == nums[i] &&
			item[1] == nums[j] &&
			item[2] == nums[k] {
			return true
		}
	}
	return false
}
