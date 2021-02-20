package algorithms

// time:O(n),space:O(1)
// 执行用时: 4 ms 内存消耗: 3.7 MB
func minSubArrayLenBySlidingWindow(target int, nums []int) int {
	// init window
	l, r := 0, -1
	min, sum := len(nums)+1, 0

	// sliding window
	for l < len(nums) {
		if r < len(nums)-1 && sum < target {
			r++
			sum += nums[r]
		} else {
			sum -= nums[l]
			l++
		}
		if sum >= target {
			window := r - l + 1
			if window == 1 {
				return 1
			} else {
				if min > window {
					min = window
				}
			}
		}
	}

	if min == len(nums)+1 {
		return 0
	}
	return min
}
