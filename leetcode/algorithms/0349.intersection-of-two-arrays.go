package algorithms

import "sort"

// time:O(n+n8longN), space:O(1)
//执行用时: 4 ms 内存消耗: 2.8 MB
func intersectionBySort(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	res := make([]int, 0)
	for i := 0; i < len(nums1); i++ {
		if i < len(nums1)-1 && nums1[i] == nums1[i+1] {
			continue
		}
		if BinarySearch(nums2, nums1[i]) != -1 {
			res = append(res, nums1[i])
		}
	}

	return res
}
