package algorithms

//time:O(1),space:O(1)
//执行用时: 0 ms 内存消耗: 2.2 MB
func arrangeCoinsByBinarySearch(n int) int {
	l, r, last := 0, n, 0
	for l <= r {
		mid := (l + r) / 2
		sum := (1 + mid) * mid / 2
		if sum <= n {
			last = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return last
}

//time:O(n),space:O(1)
//执行用时: 8 ms 内存消耗: 2.2 MB
func arrangeCoinsBySum(n int) int {
	sum, last := 0, 0
	for i := 1; i <= n; i++ {
		sum += i
		if sum <= n {
			last = i
		} else {
			break
		}
	}
	return last
}
