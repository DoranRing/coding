package algorithms

// 执行用时: 0 ms 内存消耗: 2.2 MB
func strStrByTwoPoint(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	for i := 0; i <= len(haystack)-len(needle); i++ {
		idx := i
		for j := 0; j < len(needle); j++ {
			if haystack[idx] != needle[j] {
				break
			}
			idx++
		}
		if idx-i == len(needle) {
			return i
		}
	}
	return -1
}
