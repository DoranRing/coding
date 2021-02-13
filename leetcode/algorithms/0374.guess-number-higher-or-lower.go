package algorithms

// time:O(1),space:O(1)
//执行用时: 0 ms 内存消耗: 1.9 MB
func guessNumber(n int) int {
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res == 1 {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}

func guess(num int) int {
	if num == 6 {
		return 0
	} else if num < 6 {
		return -1
	} else {
		return 1
	}
}
