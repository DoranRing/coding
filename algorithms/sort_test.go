package algorithms

import (
	"math/rand"
	"testing"
	"time"
)

type sortArgs struct {
	nums []int
}

type sortTests struct {
	name string
	args sortArgs
}

func getSortTests() []sortTests {
	return []sortTests{
		{"1", sortArgs{[]int{4, 2, 1, 3, 5, 0, 5}}},
		{"2", sortArgs{genSortNums(1000)}},
		{"3", sortArgs{genSortNums(1000)}},
		{"4", sortArgs{genSortNums(1000)}},
	}
}

func genSortNums(len int) []int {
	res := make([]int, len, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		val := rand.Intn(len * 10)
		res[i] = val
	}
	return res
}

func sorted(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func sortTest(t *testing.T, sorter func([]int)) {
	sortTests := getSortTests()
	for _, tt := range sortTests {
		sorter(tt.args.nums)
		if !sorted(tt.args.nums) {
			t.Errorf("%s not sorted", tt.name)
		}
	}
}

func TestBubbleSorter(t *testing.T) {
	sortTest(t, BubbleSorter)
}

func TestSelectSorter(t *testing.T) {
	sortTest(t, SelectSorter)
}

func TestInsertSorter(t *testing.T) {
	sortTest(t, InsertSorter)
}

func TestShellSorter(t *testing.T) {
	sortTest(t, ShellSorter)
}

func TestMergeSorter(t *testing.T) {
	sortTest(t, MergeSorter)
}

func TestQuickSorter(t *testing.T) {
	sortTest(t, QuickSorter)
}
