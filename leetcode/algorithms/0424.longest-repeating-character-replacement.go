package algorithms

func characterReplacementBySlidingWindow(s string, k int) int {
	bucket := make([]int, 26, 26)
	l, r, max := 0, -1, 0
	for l < len(s)-max {
		replace := canReplace(bucket, k, r-l+1)
		window := r - l + 1
		if r < len(s)-1 && replace {
			r++
			bucket[s[r]-'A']++
		} else {
			bucket[s[l]-'A']--
			l++
		}

		if replace && window > max {
			max = window
		}
	}
	return max
}

func canReplace(bucket []int, k, length int) bool {
	max := 0
	for _, count := range bucket {
		if count > max {
			max = count
		}
	}
	return max+k >= length
}
