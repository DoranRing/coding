package algorithms

// 执行用时: 0 ms 内存消耗: 1.9 MB
func firstBadVersion(n int) int {
	l, r := 0, n
	for l < r {
		mid := (l + r) / 2
		if !isBadVersion(mid) {
			l = mid + 1
		} else {
			r = mid
		}
	}

	return l
}

func isBadVersion(version int) bool {
	return version >= 4
}
