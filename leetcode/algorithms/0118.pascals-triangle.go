package algorithms

//执行用时: 0 ms 内存消耗: 2 MB
func generateByLoop(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	for i := 1; i <= numRows; i++ {
		if i == 1 {
			res = append(res, []int{1})
			continue
		}
		last := res[len(res)-1]
		cur := make([]int, 0)
		cur = append(cur, 1)
		for i := 0; i < len(last)-1; i++ {
			cur = append(cur, last[i]+last[i+1])
		}
		cur = append(cur, 1)
		res = append(res, cur)
	}

	return res
}
