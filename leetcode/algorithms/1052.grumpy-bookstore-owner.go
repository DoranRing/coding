package algorithms

//time:O(n),space:O(1)
//执行用时: 36 ms 内存消耗: 6.6 MB
func maxSatisfiedBySlidingWindow(customers []int, grumpy []int, X int) int {
	l, r, max, sum := 0, -1, 0, 0
	for r < len(customers)-1 {
		r++
		if grumpy[r] == 1 {
			sum += customers[r]
		}
		if max < sum {
			max = sum
		}
		if r-l+1 == X {
			if grumpy[l] == 1 {
				sum -= customers[l]
			}
			l++
		}
	}

	res := 0
	for idx, num := range customers {
		if grumpy[idx] == 0 {
			res += num
		}
	}

	return res + max
}
