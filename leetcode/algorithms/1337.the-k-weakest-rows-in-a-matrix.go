package algorithms

import (
	"sort"
)

type Entry struct {
	Idx int

	Val int
}

//time:O(n*logN),space:O(1)
//执行用时: 12 ms 内存消耗: 5 MB
func kWeakestRows(mat [][]int, k int) []int {
	entries := make([]*Entry, 0, len(mat))
	for idx, arr := range mat {
		l, r, last := 0, len(arr)-1, -1
		for l <= r {
			mid := (l + r) / 2
			if arr[mid] == 0 {
				last = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
		if last == -1 {
			entries = append(entries, &Entry{idx, len(arr)})
		} else {
			entries = append(entries, &Entry{idx, last - 0})
		}
	}
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Val < entries[j].Val
	})

	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, entries[i].Idx)
	}
	return res
}
