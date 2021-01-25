package algorithms

import "sort"

// 使用哈希表记录第一个集合的元素以及次数，然后比对第二个集合的元素。
// time: O(n) space: O(n)
// 执行用时: 4 ms 内存消耗: 3.1 MB
func intersectByHash(nums1 []int, nums2 []int) []int {
	hash := make(map[int]int)
	var short []int
	var longer []int
	if len(nums1) < len(nums2) {
		short = nums1
		longer = nums2
	} else {
		short = nums2
		longer = nums1
	}

	for _, v := range short {
		var count int
		if exist, ok := hash[v]; ok {
			count = exist + 1
		} else {
			count = 1
		}
		hash[v] = count
	}

	res := make([]int, 0)
	for _, v := range longer {
		if exist, ok := hash[v]; ok && exist > 0 {
			res = append(res, v)
			hash[v] = exist - 1
		}
	}

	return res
}

// 先排序，然后一次比对两个集合，相等则加入结果，A大于B则B指针后移，反之A指针后移
// time: O(n+logN) space: O(0)
// 执行用时: 4 ms 内存消耗: 2.8 MB
func intersectByTwoPoint(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	res := make([]int, 0)
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}

	return res
}
