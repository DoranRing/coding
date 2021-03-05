package algorithms

// time:O(n),space:O(1)
// 执行用时: 16 ms 内存消耗: 6.3 MB
func sortArrayByParityIIByTwoPoint(A []int) []int {
	for even, odd := 0, 1; even < len(A); {
		if A[even]%2 != 0 {
			A[even], A[odd] = A[odd], A[even]
			odd += 2
		} else {
			even += 2
		}
	}
	return A
}
