package algorithms

// 通过双指针分别指向集合的头和尾，依次交换两个指针的值
// time: O(n), space:O(1)
//执行用时: 44 ms 内存消耗: 6.5 MB
func reverseStringByPoint(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
