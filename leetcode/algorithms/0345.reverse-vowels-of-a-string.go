package algorithms

// 通过对撞指针
//time:O(n),space:O(1)
func reverseVowelsByCollisionPointer(s string) string {
	data := make([]uint8, len(s), len(s))
	for l, r := 0, len(s)-1; l <= r; {
		if !isVowels(s[l]) {
			data[l] = s[l]
			l++
		} else if !isVowels(s[r]) {
			data[r] = s[r]
			r--
		} else {
			data[l] = s[r]
			data[r] = s[l]
			l, r = l+1, r-1
		}
	}

	return string(data)
}

func isVowels(c uint8) bool {
	return c == 'a' || c == 'A' ||
		c == 'e' || c == 'E' ||
		c == 'i' || c == 'I' ||
		c == 'o' || c == 'O' ||
		c == 'u' || c == 'U'
}
