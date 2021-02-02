package algorithms

// 执行用时: 0 ms 内存消耗: 1.9 MB
func firstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if isBadVersion(mid) && !isBadVersion(mid-1) {
			return mid
		} else if isBadVersion(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

func isBadVersion(version int) bool {
	return version >= 4
}
