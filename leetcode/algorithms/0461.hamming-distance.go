package algorithms

// 执行用时: 0 ms 内存消耗: 2 MB
func hammingDistanceByMath(x int, y int) int {
	sum := 0
	for x != 0 || y != 0 {
		xBit, yBit := x%2, y%2
		if xBit > yBit {
			sum += xBit - yBit
		} else {
			sum += yBit - xBit
		}
		x, y = x/2, y/2
	}

	return sum
}

func hammingDistanceByBit(x int, y int) int {
	sum := 0
	for x != 0 || y != 0 {
		xBit, yBit := x&1, y&1
		if xBit > yBit {
			sum += xBit - yBit
		} else {
			sum += yBit - xBit
		}
		x, y = x>>1, y>>1
	}

	return sum
}
