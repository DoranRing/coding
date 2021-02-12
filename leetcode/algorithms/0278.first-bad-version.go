package algorithms

// 执行用时: 0 ms 内存消耗: 1.9 MB
func firstBadVersion(n int) int {
	l, r, ans := 0, n, -1
	for l <= r {
		mid := (l + r) / 2
		if isBadVersion(mid) {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func isBadVersion(version int) bool {
	return version >= 4
}
