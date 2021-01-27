package algorithms

// 双指针分别指向字符串的首尾，如果指针指向的字符不是要求的字符则跳过，依次比较指针指向的字符是否相等，
// 执行用时: 0 ms 内存消耗: 3.2 MB
func isPalindromeByTwoPoint(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; {
		if !isWord(runes, i) {
			i++
			continue
		}
		if !isWord(runes, j) {
			j--
			continue
		}
		if !equal(runes, i, j) {
			return false
		}
		i++
		j--
	}

	return true
}

func isWord(runes []rune, i int) bool {
	val := runes[i]
	return (val >= '0' && val <= '9') ||
		(val >= 'a' && val <= 'z') ||
		(val >= 'A' && val <= 'Z')
}

func equal(runes []rune, i, j int) bool {
	val1, val2 := runes[i], runes[j]
	// A-Z to a-z
	if val1 >= 'A' && val1 <= 'Z' {
		val1 += 'a' - 'A'
	}
	if val2 >= 'A' && val2 <= 'Z' {
		val2 += 'a' - 'A'
	}

	return val1 == val2
}
