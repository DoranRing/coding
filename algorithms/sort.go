package algorithms

// BubbleSorter 冒泡排序(稳定排序)
// time:O(n2), space:O(1)
func BubbleSorter(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// SelectSorter 选择排序(稳定排序)
// time:O(n2), space:O(1)
func SelectSorter(nums []int) {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
}

// InsertSorter 插入排序(稳定排序)
// time:O(n2), space:O(1)
func InsertSorter(nums []int) {
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0; j-- {
			if nums[j] >= nums[j-1] {
				break
			}
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
}

// ShellSorter 希尔排序(不稳定排序)
// time:O(n1.3-2), space:O(1)
func ShellSorter(nums []int) {
	step := len(nums) / 2
	for step > 0 {
		// 每组分别排序
		for i := 0; i < step; i++ {
			// 插入排序
			for j := i; j < len(nums); j += step {
				for k := j; k >= step; k -= step {
					if nums[k] >= nums[k-step] {
						break
					}
					nums[k], nums[k-step] = nums[k-step], nums[k]
				}
			}
		}
		step /= 2
	}
}

// MergeSorter 归并排序(稳定排序)
// time:O(n*logN), space:O(n)
func MergeSorter(nums []int) {
	if len(nums) <= 1 {
		return
	}
	temp := make([]int, len(nums), len(nums))
	mergeSorter(nums, temp, 0, len(nums)-1)
}

func mergeSorter(nums, temp []int, left, right int) {
	if left == right {
		// 一个元素
		return
	} else if left+1 == right {
		// 排序两个相邻元素
		merge(nums, temp, left, right, right)
	} else {
		mid := (left + right) / 2
		// 分治前半部分
		mergeSorter(nums, temp, left, mid)
		// 分支后半部分
		mergeSorter(nums, temp, mid+1, right)
		// 归并前、后半部分
		merge(nums, temp, left, mid+1, right)
	}
}

func merge(nums, temp []int, l1, l2, r2 int) {
	idx := l1
	for i, j := l1, l2; i < l2 || j <= r2; idx++ {
		if i < l2 && j <= r2 {
			if nums[i] < nums[j] {
				temp[idx] = nums[i]
				i++
			} else {
				temp[idx] = nums[j]
				j++
			}
		} else if i < l2 {
			temp[idx] = nums[i]
			i++
		} else if j <= r2 {
			temp[idx] = nums[j]
			j++
		}
	}
	for i := l1; i <= r2; i++ {
		nums[i] = temp[i]
	}
}

// QuickSorter 快速排序
// time:O(n*longN), space:O(1)
func QuickSorter(nums []int) {
	if len(nums) <= 1 {
		return
	}
	quickSorter(nums, 0, len(nums)-1)
}

func quickSorter(nums []int, left, right int) {
	if left >= right {
		return
	}

	refVal, refIdx := nums[left], left
	for i, j := 0, len(nums)-1; i <= j; {
		// 小于值向前移动
		for j >= refIdx && nums[j] >= refVal {
			j--
		}
		if j >= refIdx {
			nums[refIdx] = nums[j]
			refIdx = j
		}
		// 大于值向后移动
		for i <= refIdx && nums[i] <= refVal {
			i++
		}
		if i <= refIdx {
			nums[refIdx] = nums[i]
			refIdx = i
		}
	}

	nums[refIdx] = refVal
	quickSorter(nums, left, refIdx-1)
	quickSorter(nums, refIdx+1, right)
}

// BucketSorter 桶排序(稳定排序)
// time:O(m+n),space:O(m)
func BucketSorter(nums []int, max int) {
	bucket := make([]int, max+1, max+1)
	for _, num := range nums {
		bucket[num]++
	}

	idx := 0
	for i := 0; i < len(bucket); i++ {
		for bucket[i] > 0 {
			nums[idx] = i
			idx++
			bucket[i]--
		}
	}
}
