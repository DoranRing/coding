package algorithms

// 执行用时: 0 ms 内存消耗: 2.3 MB
func mergeByTwoPoint(nums1 []int, m int, nums2 []int, n int) {
	idx := m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[idx] = nums1[m]
			m--
		} else {
			nums1[idx] = nums2[n]
			n--
		}
		idx--
	}

	for ; m >= 0; m-- {
		nums1[idx] = nums1[m]
		idx--
	}
	for ; n >= 0; n-- {
		nums1[idx] = nums2[n]
		idx--
	}

}
