package algorithms

//time:O(n),space:O(n)
//执行用时: 16 ms 内存消耗: 6.9 MB
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) <= 0 || r*c > len(nums)*len(nums[0]) {
		return nums
	}
	res := make([][]int, r, r)
	count := 0
	for _, num := range nums {
		for _, v := range num {
			if res[count/c] == nil {
				res[count/c] = make([]int, c, c)
			}
			res[count/c][count%c] = v
			count++
		}
	}
	return res
}
