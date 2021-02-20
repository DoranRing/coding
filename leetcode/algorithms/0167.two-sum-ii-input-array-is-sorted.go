package algorithms

// 利用序列数据特性，使用二分查找
// time:O(n), space:O(1)
// 执行用时: 4 ms 内存消耗: 2.9 MB
func twoSumByBinarySearch(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		num, left, right := target-numbers[i], i+1, len(numbers)-1
		for left <= right {
			mid := (left + right) / 2
			if numbers[mid] == num {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > num {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return []int{}
}

// 通过碰撞指针
// time:O(n), space:O(1)
// 执行用时: 4 ms 内存消耗: 2.9 MB
func twoSumByCollisionPointer(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for l < r {
		sum := numbers[l] + numbers[r]
		if sum == target {
			return []int{l + 1, r + 1}
		} else if sum < target {
			l++
		} else {
			r--
		}
	}
	return []int{}
}
