package algorithms

// 执行用时: 0 ms 内存消耗: 1.9 MB
func hammingWeightByMath(num uint32) int {
	count := 0
	for num != 0 {
		if num%2 == 1 {
			count++
		}
		num = num / 2
	}

	return count
}

// 执行用时: 0 ms 内存消耗: 1.9 MB
func hammingWeightByBit(num uint32) int {
	count := uint32(0)
	for num != 0 {
		count += num & 1
		num = num >> 1
	}
	return int(count)
}
