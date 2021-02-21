package algorithms

//time:O(n),space:O(128)
//执行用时: 4 ms 内存消耗: 2.5 MB
func lengthOfLongestSubstringBySlidingWindow(s string) int {
	bucket := make([]int, 128, 128)
	l, r := 0, -1
	max := 0
	for r < len(s)-1 {
		if r == -1 || bucket[s[r+1]] == 0 {
			r++
			bucket[s[r]]++
			if max < (r - l + 1) {
				max = r - l + 1
			}
		} else {
			bucket[s[l]]--
			l++
		}
	}
	return max
}
