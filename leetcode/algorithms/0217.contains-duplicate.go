package algorithms

import "sort"

// 通过哈希表记录元素的出现情况

// time: O(n) space:O(n)
// 执行用时：24 ms 内存消耗： 8.1 MB
func containsDuplicateByHash(nums []int) bool {
	hash := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := hash[nums[i]]; ok {
			return true
		}
		hash[nums[i]] = true
	}
	return false
}

// 先排序数组，然后比较相邻元素相等

// time: O(logN) space: O(1)
// 执行用时：24 ms 内存消耗： 6 MB
func containsDuplicateBySort(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}
