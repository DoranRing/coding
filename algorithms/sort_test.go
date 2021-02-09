package algorithms

import "testing"

type sortArgs struct {
	nums []int
}

type sortTests struct {
	name string
	args sortArgs
}

func getSortTests() []sortTests {
	return []sortTests{
		{"1", sortArgs{[]int{1, 2, 3, 4}}},
		{"2", sortArgs{[]int{4, 3, 2, 1}}},
		{"2", sortArgs{[]int{2, 3, 4, 1}}},
		{"3", sortArgs{[]int{}}},
	}
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
