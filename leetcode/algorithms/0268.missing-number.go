package algorithms

// 执行用时: 16 ms 内存消耗: 6.1 MB
func missingNumberByMath(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}

	should := (1 + len(nums)) * len(nums) / 2
	return should - sum
}

// 执行用时: 16 ms 内存消耗: 6.1 MB
func missingNumberByXOR(nums []int) int {
	xor := 0
	for i := 0; i <= len(nums); i++ {
		xor ^= i
	}
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
