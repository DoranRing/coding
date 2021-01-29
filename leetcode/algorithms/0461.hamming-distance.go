package algorithms

// 执行用时: 0 ms 内存消耗: 2 MB
func hammingDistance(x int, y int) int {
	sum := 0
	for x != 0 || y != 0 {
		xBit := x % 2
		yBit := y % 2
		if xBit > yBit {
			sum += xBit - yBit
		} else {
			sum += yBit - xBit
		}
		x = x / 2
		y = y / 2
	}

	return sum
}
