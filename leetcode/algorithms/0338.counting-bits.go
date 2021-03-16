package algorithms

func countBitsByCompute(num int) []int {
	res := make([]int, num+1)
	for i := 0; i <= num; i++ {
		cur, sum := i, 0
		for cur > 0 {
			sum = sum + cur%2
			cur = cur >> 1
		}
		res[i] = sum
	}
	return res
}
