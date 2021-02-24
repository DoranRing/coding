package algorithms

// time:O(n),space:O(1)
//执行用时: 4 ms 内存消耗: 2.9 MB
func flipAndInvertImage(A [][]int) [][]int {
	for _, nums := range A {
		for i, j := 0, len(nums)-1; i <= j; i, j = i+1, j-1 {
			if nums[i] == nums[j] {
				val := 0
				if nums[i] == 0 {
					val = 1
				}
				nums[i] = val
				nums[j] = val
			}
		}
	}
	return A
}
