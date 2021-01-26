package algorithms

// 通过哈希表记录出现过的元素以及下标，后面元素在哈希表中寻找符合要求的特定元素，找到则返回结果
// time:O(n) space:O(n)
// 执行用时: 4 ms 内存消耗: 3 MB
func twoSumByHash(nums []int, target int) []int {
	hash := make(map[int]int)
	for idx, num := range nums {
		if val, ok := hash[target-num]; ok {
			return []int{val, idx}
		} else {
			hash[num] = idx
		}
	}

	return nil
}

// 排序方式不适用本题，本题的结果需要是元素的下标，排序覆盖元素应有的下标
//func twoSumBySort(nums []int, target int) []int {
//	sort.Ints(nums)
//	for i := 0; i < len(nums); i++ {
//		var needIdx = binarySearch(nums, i+1, len(nums), target-nums[i])
//		if needIdx != -1 {
//			return []int{i, needIdx}
//		}
//	}
//	return nil
//}
//
//func binarySearch(nums []int, l, r, val int) int {
//	for l <= r {
//		mid := (l + r) / 2
//		if nums[mid] == val {
//			return mid
//		} else if nums[mid] > val {
//			r = mid - 1
//		} else {
//			l = mid + 1
//		}
//	}
//	return -1
//}
