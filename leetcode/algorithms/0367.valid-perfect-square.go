package algorithms

// time:O(1), space:O(1)
//执行用时: 0 ms 内存消耗: 1.9 MB
func isPerfectSquare(num int) bool {
	num64 := int64(num)
	var l, r int64 = 0, 50000
	for l <= r {
		mid := (l + r) / 2
		sq := mid * mid
		if sq == num64 {
			return true
		} else if sq < num64 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
