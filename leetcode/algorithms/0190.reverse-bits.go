package algorithms

// 执行用时: 4 ms 内存消耗: 2.6 MB
func reverseBitsByMath(num uint32) uint32 {
	var res uint32 = 0
	var step uint32 = 1 << 31
	for num != 0 {
		res += (num % 2) * step
		step = step >> 1
		num /= 2
	}

	return res
}

// 执行用时: 4 ms 内存消耗: 2.6 MB
func reverseBitsByBit(num uint32) uint32 {
	res := uint32(0)
	step := uint32(31)
	for num != 0 {
		// new result, don't change num value
		res += (num & 1) << step
		num = num >> 1
		step -= 1
	}
	return res
}
