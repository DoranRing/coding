package algorithms

// 执行用时: 36 ms 内存消耗: 6.3 MB
func isPowerOfThreeByLoop(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n /= 3
	}
	return n == 1
}
