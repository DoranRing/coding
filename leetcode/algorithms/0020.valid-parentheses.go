package algorithms

// 执行用时: 0 ms 内存消耗: 2 MB
func isValidByStack(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		// match, ']' - '[' = 2, ')'-'("=1
		if len(stack) > 0 &&
			char-stack[len(stack)-1] < 3 &&
			char-stack[len(stack)-1] > 0 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
