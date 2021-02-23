package algorithms

//time:O(n),space:O(1)
//执行用时: 68 ms 内存消耗: 6.9 MB
func longestOnesBySlidingWindow(A []int, K int) int {
	l, r, max, maxCount := 0, -1, 0, 0

	for r < len(A)-1 {
		r++
		if A[r] == 1 {
			maxCount++
		}
		if r-l+1 > maxCount+K {
			if r-l > max {
				max = r - l
			}
			if A[l] == 1 {
				maxCount--
			}
			l++
		} else if r == len(A)-1 {
			if r-l+1 > max {
				max = r - l + 1
			}
		}
	}

	return max
}
