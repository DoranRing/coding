package algorithms

// time:O(1),space:O(1)
//执行用时: 16 ms 内存消耗: 6.3 MB
func countNegatives(grid [][]int) int {
	sum := 0
	for i := len(grid) - 1; i >= 0; i-- {
		l, r, idx := 0, len(grid[i])-1, -1
		for l <= r {
			mid := (l + r) / 2
			if grid[i][mid] < 0 {
				idx = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if idx == -1 {
			break
		}
		sum += len(grid[i]) - idx
	}
	return sum
}
