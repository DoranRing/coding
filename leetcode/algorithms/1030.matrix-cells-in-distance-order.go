package algorithms

import (
	"sort"
)

type SumIntSlice struct {
	slice []int
	sum   int
}

// time:O(RC+RC*logN),space:O(RC)
// 执行用时: 20 ms 内存消耗: 7.5 MB
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	sumRes := make([]*SumIntSlice, 0, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			diffR := i - r0
			diffC := j - c0
			if diffR < 0 {
				diffR = -diffR
			}
			if diffC < 0 {
				diffC = -diffC
			}
			sum := diffC + diffR
			sumRes = append(sumRes, &SumIntSlice{[]int{i, j}, sum})
		}
	}

	// sort
	sort.Slice(sumRes, func(i, j int) bool {
		return sumRes[i].sum < sumRes[j].sum
	})

	res := make([][]int, 0, len(sumRes))
	for i := 0; i < len(sumRes); i++ {
		res = append(res, sumRes[i].slice)
	}
	return res
}

//time:O(RC), space:O(RC)
// 执行用时: 16 ms 内存消耗: 7.2 MB
func allCellsDistOrderByBucket(R int, C int, r0 int, c0 int) [][]int {
	bucket := make([][][]int, R*C)
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			diffR := i - r0
			diffC := j - c0
			if diffR < 0 {
				diffR = -diffR
			}
			if diffC < 0 {
				diffC = -diffC
			}
			sum := diffC + diffR
			arr := bucket[sum]
			if bucket[sum] == nil {
				arr = make([][]int, 0)
			}
			bucket[sum] = append(arr, []int{i, j})
		}
	}

	res := make([][]int, 0, R*C)
	for i := 0; i < len(bucket); i++ {
		if bucket[i] == nil {
			continue
		}
		for j := 0; j < len(bucket[i]); j++ {
			res = append(res, bucket[i][j])
		}
	}

	return res
}
