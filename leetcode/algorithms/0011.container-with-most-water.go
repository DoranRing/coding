package algorithms

// time:O(n2),space:O(1)
//执行用时: 928 ms 内存消耗: 6.3 MB
func maxAreaByForce(height []int) int {
	max := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			min := height[i]
			if min > height[j] {
				min = height[j]
			}
			area := (j - i) * min
			if area > max {
				max = area
			}
		}
	}
	return max
}

// time:O(n),space:O(1)
// 执行用时: 20 ms 内存消耗: 6.3 MB
func maxAreaByCollisionPointer(height []int) int {
	max := 0
	l, r := 0, len(height)-1
	for l < r {
		min := height[l]
		if min > height[r] {
			min = height[r]
		}
		area := (r - l) * min
		if area > max {
			max = area
		}

		if min == height[l] {
			l++
		} else {
			r--
		}
	}
	return max
}
