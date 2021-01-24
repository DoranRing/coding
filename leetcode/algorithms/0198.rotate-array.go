package algorithms

// 使用缓存空间
// |123|4|  -> |4|123|

// time: O(n) space(n)
// 执行用时: 0 ms  内存消耗: 3.2 MB
func rotateByBuffer(nums []int, k int) {
	k = k % len(nums)
	buf := make([]int, k)
	endIdx := len(nums) - k - 1
	copy(buf, nums[endIdx+1:])
	for i := endIdx; i >= 0; i-- {
		nums[i+k] = nums[i]
	}
	for i := 0; i < k; i++ {
		nums[i] = buf[i]
	}
}

// 通过三次翻转
// |123|45| -> |54|321|(第一次全部翻转) -> |45|321|(翻转前部分) -> |45|123|(翻转后部分)

// time: O(n) space(1)
// 执行用时: 4 ms 内存消耗: 3.1 MB
func rotateByReverse(nums []int, k int) {
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

func reverse(nums []int, l, r int) {
	for ; l < r; l, r = l+1, r-1 {
		nums[l], nums[r] = nums[r], nums[l]
	}
}
