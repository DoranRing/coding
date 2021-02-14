package algorithms

// time:O(1),space:O(1)
//执行用时: 0 ms 内存消耗: 2.6 MB
func nextGreatestLetter(letters []byte, target byte) byte {
	l, r := 0, len(letters)-1
	var res byte = 0
	for l <= r {
		mid := (l + r) / 2
		if letters[mid] > target {
			res = letters[mid]
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if res == 0 && len(letters) > 0 {
		return letters[0]
	}
	return res
}
