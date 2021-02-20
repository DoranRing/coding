package algorithms

// BinarySearchByBase 二分查找
// 匹配多个时选取的值不固定
func BinarySearchByBase(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}

// BinarySearchByRangeFirst 二分查找
// 查询符合条件的第一个元素
func BinarySearchByRangeFirst(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)
	for l < r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	if l != len(nums) && nums[l] == target {
		return l
	}
	return -1
}

// CollisionPointer 碰撞指针
// 集合中两个元素的关联性，这里找出两个元素和等于target
func CollisionPointer(numbers []int, target int) []int {
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

// SlowFastPoint 快慢指针
// 集合中两个元素的关联性，这里找出尾部节点是否指向了首部节点
func SlowFastPoint(head *ListNode) bool {
	var fast, slow *ListNode
	for head != nil {
		// fast+1
		if fast == nil {
			fast = head
		} else {
			fast = fast.Next
		}

		// fast+1,slow+1
		if fast != nil && fast.Next != nil {
			fast = fast.Next
			if slow == nil {
				slow = head
			} else {
				slow = slow.Next
			}
		} else {
			// next is nil
			return false
		}

		// equal
		if slow == fast {
			return true
		}
	}

	return false
}

type ListNode struct {
	Val int

	Next *ListNode
}

// SlidingWindow 滑动窗口
// 检查一个子集符合要求
func SlidingWindow(target int, nums []int) int {
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
