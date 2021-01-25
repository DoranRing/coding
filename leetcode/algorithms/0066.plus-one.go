package algorithms

//执行用时: 0 ms 内存消耗: 2 MB
func plusOneByLoop(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		r := digits[i] + 1
		if r == 10 {
			digits[i] = 0
		} else {
			digits[i] = r
			return digits
		}
	}

	res := make([]int, len(digits)+1)
	res[0] = 1
	return res
}
