package algorithms

//执行用时: 16 ms 内存消耗: 5.4 MB
func findMedianSortedArraysByMergeTraversal(nums1 []int, nums2 []int) float64 {
	// 确定中位数位置
	total := len(nums1) + len(nums2)
	idx1, idx2, count := -1, -1, 0
	if total%2 == 0 {
		idx1 = total/2 - 1
		idx2 = total / 2
	} else {
		idx1 = total / 2
	}

	// 归并排序的方式遍历，遍历到中位数位置时取值
	nums := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		var min int
		if nums1[i] < nums2[j] {
			min = nums1[i]
			i++
		} else {
			min = nums2[j]
			j++
		}

		if count == idx1 || count == idx2 {
			nums = append(nums, min)
			if idx2 == -1 || len(nums) == 2 {
				return mapResult(nums)
			}
		}
		count++
	}
	for ; i < len(nums1); i++ {
		if count == idx1 || count == idx2 {
			nums = append(nums, nums1[i])
			if idx2 == -1 || len(nums) == 2 {
				return mapResult(nums)
			}
		}
		count++
	}
	for ; j < len(nums2); j++ {
		if count == idx1 || count == idx2 {
			nums = append(nums, nums2[j])
			if idx2 == -1 || len(nums) == 2 {
				return mapResult(nums)
			}
		}
		count++
	}

	return 0
}

func mapResult(nums []int) float64 {
	// 计算结果
	if len(nums) == 1 {
		return float64(nums[0])
	} else if len(nums) == 2 {
		return (float64(nums[0]) + float64(nums[1])) / 2
	}
	return 0
}
