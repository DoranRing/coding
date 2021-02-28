package algorithms

import (
	"strconv"
)

//time:O(s),space:O(1)
// 执行用时: 0 ms 内存消耗: 1.9 MB
func summaryRanges(nums []int) []string {
	res := make([]string, 0)
	idx := 0
	for i := 0; i < len(nums); i++ {
		if i >= len(nums)-1 || nums[i+1]-nums[i] != 1 {
			var item string
			if i == idx {
				item = strconv.Itoa(nums[i])
			} else {
				item = strconv.Itoa(nums[idx]) + "->" + strconv.Itoa(nums[i])
			}
			res = append(res, item)
			idx = i + 1
		}
	}
	return res
}
