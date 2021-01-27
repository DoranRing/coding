package algorithms

import "math"

//执行用时: 0 ms 内存消耗: 2.3 MB
func myAtoi(s string) int {
	idx := 0
	for i := 0; i < len(s); i++ {
		if s[idx] != ' ' {
			break
		}
		idx++
	}

	symbol := int64(1)
	if idx < len(s) && s[idx] == '-' {
		symbol = -1
		idx++
	} else if idx < len(s) && s[idx] == '+' {
		idx++
	}

	var res int64
	for ; idx < len(s); idx++ {
		if s[idx] < '0' || s[idx] > '9' {
			break
		}
		res = res*10 + int64(s[idx]-'0')*symbol
		if symbol == int64(1) && res >= math.MaxInt32 {
			return math.MaxInt32
		}
		if symbol == int64(-1) && res <= math.MinInt32 {
			return math.MinInt32
		}
	}

	return int(res)
}
