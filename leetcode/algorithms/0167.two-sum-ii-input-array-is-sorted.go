package algorithms

// 执行用时: 4 ms 内存消耗: 2.9 MB
// time:O(n), space:O(1)
func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		idx := twoSumByBinarySearch(numbers, target-numbers[i], i+1, len(numbers)-1)
		if idx != -1 {
			return []int{i + 1, idx + 1}
		}
	}
	return []int{}
}

func twoSumByBinarySearch(nums []int, num, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == num {
			return mid
		} else if nums[mid] > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
