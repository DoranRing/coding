package algorithms

import "math"

// 执行用时: 0 ms 内存消耗: 2.2 MB
func reverseByMath(x int) int {
	var res int64
	positive := false
	if x > 0 {
		positive = true
	}
	for x != 0 {
		res = res*10 + int64(x%10)
		x /= 10
	}

	if positive && res >= math.MaxInt32 {
		return 0
	}
	if !positive && res <= math.MinInt32 {
		return 0
	}

	return int(res)
}
