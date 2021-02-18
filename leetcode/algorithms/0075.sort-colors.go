package algorithms

// time:O(n),space:(3)
// 执行用时: 0 ms 内存消耗: 2 MB
func sortColorsByBucket(nums []int) {
	bucket := make([]int, 3, 3)
	for _, num := range nums {
		bucket[num]++
	}

	count := 0
	for idx, item := range bucket {
		for i := 0; i < item; count, i = count+1, i+1 {
			nums[count] = idx
		}
	}
}
