package algorithms

// time:O(n),space:O(1)
// 执行用时: 4 ms 内存消耗: 2.8 MB
func equalSubstringBySlidingWindow(s string, t string, maxCost int) int {
	l, r, max, sumConst := 0, -1, 0, 0

	for r < len(s)-1 {
		if maxCost >= sumConst {
			r++
			diff := int(s[r]) - int(t[r])
			if diff < 0 {
				diff = 0 - diff
			}
			sumConst += diff
		} else {
			diff := int(s[l]) - int(t[l])
			if diff < 0 {
				diff = 0 - diff
			}
			l++
			sumConst -= diff
		}

		if maxCost >= sumConst && r-l+1 > max {
			max = r - l + 1
		}
	}

	return max
}
