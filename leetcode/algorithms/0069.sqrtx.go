package algorithms

// 执行用时: 4 ms 内存消耗: 2.2 MB
// time:O(logN),space:O(1)
func mySqrt(x int) int {
	left, right, res := 0, x, -1
	for left <= right {
		mid := (left + right) / 2
		square := mid * mid
		if square <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}
