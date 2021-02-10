package algorithms

// BinarySearch 二分查找
func BinarySearch(sortNums []int, val int) int {
	return binarySearch(sortNums, val, 0, len(sortNums)-1)
}

func binarySearch(nums []int, val, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == val {
			return mid
		} else if nums[mid] > val {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
