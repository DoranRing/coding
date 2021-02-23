package algorithms

// time:O(n), space:(1)
// 执行用时: 76 ms 内存消耗: 7.2 MB
func maxTurbulenceSizeBySlidingWindow(arr []int) int {
	l, r, max, pre := 0, 0, 1, -1
	for r < len(arr)-1 {
		// 当前元素与下一个元素的符号
		bit := 0
		if arr[r] > arr[r+1] {
			bit = 1
		} else if arr[r] == arr[r+1] {
			bit = -1
			pre = -1
		}

		if bit^pre == 0 {
			// 当前符号与左右两个元素的符号一样
			if r-l+1 > max {
				max = r - l + 1
			}
			if bit == -1 {
				l = r + 1
			} else {
				l = r
			}
		} else if r+1 == len(arr)-1 {
			// 符合且到达了末尾边界
			if r-l+2 > max {
				max = r - l + 2
			}
		}

		r++
		pre = bit
	}
	return max
}
