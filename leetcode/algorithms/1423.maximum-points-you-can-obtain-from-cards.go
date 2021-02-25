package algorithms

func maxScoreBySlidingWindow(cardPoints []int, k int) int {
	l, r, max, sum := -k, -k-1, 0, 0

	for r < k-1 {
		r++
		rIdx := r
		if r < 0 {
			rIdx = len(cardPoints) + r
		}
		sum += cardPoints[rIdx]

		if r-l+1 > k {
			lIdx := l
			if l < 0 {
				lIdx = len(cardPoints) + l
			}
			sum -= cardPoints[lIdx]
			l++
		}

		if sum > max {
			max = sum
		}
	}

	return max
}
