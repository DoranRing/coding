package algorithms

// BinarySearch 二分查找
func BinarySearch(nums []int, val int) int {
	left, right := 0, len(nums)-1
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
