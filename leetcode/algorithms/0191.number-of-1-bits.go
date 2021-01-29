package algorithms

// 执行用时: 0 ms 内存消耗: 1.9 MB
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		if num%2 == 1 {
			count++
		}
		num = num / 2
	}

	return count
}
