package algorithms

// time:O(1),space:O(1)
//执行用时: 8 ms 内存消耗: 4.5 MB
func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2
		if arr[mid] < arr[mid+1] {
			l = mid + 1
		} else if arr[mid] < arr[mid-1] {
			r = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
