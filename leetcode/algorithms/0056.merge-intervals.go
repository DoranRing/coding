package algorithms

import "sort"

func mergeBySort(intervals [][]int) [][]int {
	// 空值处理
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 排序
	var sorter IntSliceSorter = intervals
	sort.Sort(sorter)

	// 合并
	var res [][]int
	var point []int = nil
	for _, col := range intervals {
		if point == nil {
			point = col
		} else {
			if point[1] < col[0] {
				res = append(res, point)
				point = col
			} else {
				point = mergeOverLapping(point, col)
			}
		}
	}
	if point != nil {
		res = append(res, point)
	}
	return res
}

func mergeOverLapping(a, b []int) []int {
	max := b[1]
	if a[1] > max {
		max = a[1]
	}
	return []int{a[0], max}
}

type IntSliceSorter [][]int

func (s IntSliceSorter) Len() int {
	return len(s)
}

func (s IntSliceSorter) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s IntSliceSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
