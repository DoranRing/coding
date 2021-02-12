package algorithms

import (
	"fmt"
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

func getSortTests(len, max int) []sortTests {
	return []sortTests{
		//{"1", sortArgs{[]int{4, 2, 1, 3, 5, 0, 5}}},
		{"2", sortArgs{genIntNums(len, max)}},
		{"3", sortArgs{genIntNums(len, max)}},
		{"4", sortArgs{genIntNums(len, max)}},
	}
}

func genIntNums(len, max int) []int {
	res := make([]int, len, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		val := rand.Intn(max)
		res[i] = val
	}
	return res
}

func genInt32Nums(len int, max int32) []int32 {
	res := make([]int32, len, len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		val := rand.Int31n(max)
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

func sortedInt32(nums []int32) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return false
		}
	}
	return true
}

func sortTest(t *testing.T, sorter func([]int), len, max int) {
	sortTests := getSortTests(len, max)
	for _, tt := range sortTests {
		sorter(tt.args.nums)
		if !sorted(tt.args.nums) {
			t.Errorf("%s not sorted", tt.name)
		}
	}
}

func benchmarkTest(sorter func([]int), len, max int) {
	sortTests := getSortTests(len, max)
	for _, tt := range sortTests {
		sorter(tt.args.nums)
	}
}

func TestBubbleSorter(t *testing.T) {
	sortTest(t, BubbleSorter, 1000, 1000*10)
}

//2439192 ns/op
func BenchmarkBubbleSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTest(BubbleSorter, 1000, 1000*10)
	}
}

func TestSelectSorter(t *testing.T) {
	sortTest(t, SelectSorter, 1000, 1000*10)
}

// 1370888 ns/op
func BenchmarkSelectSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTest(SelectSorter, 1000, 1000*10)
	}
}

func TestInsertSorter(t *testing.T) {
	sortTest(t, InsertSorter, 1000, 1000*10)
}

// 995473 ns/op
func BenchmarkInsertSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTest(InsertSorter, 1000, 1000*10)
	}
}

func TestShellSorter(t *testing.T) {
	sortTest(t, ShellSorter, 1000, 1000*10)
}

// 564152 ns/op
func BenchmarkShellSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTest(ShellSorter, 1000, 1000*10)
	}
}

func TestMergeSorter(t *testing.T) {
	sortTest(t, MergeSorter, 1000, 1000*10)
}

// 267570 ns/op
func BenchmarkMergeSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTest(MergeSorter, 1000, 1000*10)
	}
}

func TestQuickSorter(t *testing.T) {
	sortTest(t, QuickSorter, 1000, 1000*10)
}

// 261440-1015922 ns/op
func BenchmarkQuickSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTest(QuickSorter, 1000, 1000*10)
	}
}

func TestBucketSorter(t *testing.T) {
	sortTest(
		t,
		func(nums []int) {
			BucketSorter(nums, 500)
		},
		1000, 500)
}

// 101836 ns/op
func BenchmarkBucketSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchmarkTest(
			func(nums []int) {
				BucketSorter(nums, 500)
			},
			1000, 500,
		)
	}
}

func TestHardList(t *testing.T) {
	hardList := DefaultHardList()
	defer hardList.Close()
	var num1 int32 = 20
	var num2 int32 = 21
	hardList.Append(num1)
	hardList.Append(num2)
	val1 := hardList.Get(0)
	val2 := hardList.Get(1)
	if num1 != val1 {
		t.Errorf("should: %d, actual: %d\n", num1, val1)
	}
	if num2 != val2 {
		t.Errorf("should: %d, actual: %d\n", num2, val2)
	}

	var num3 int32 = 30
	hardList.Set(1, num3)
	val3 := hardList.Get(1)
	if num3 != val3 {
		t.Errorf("should: %d, actual: %d\n", num3, val3)
	}

	nums := hardList.All()
	for _, num := range nums {
		fmt.Printf("%d, ", num)
	}
}

func TestHardBubbleSorterValid(t *testing.T) {
	nums := genInt32Nums(1000, 1000*100)
	sortedNums := HardBubbleSorterValid(nums)
	if !sortedInt32(sortedNums) {
		t.Errorf("not sorted")
	}
}

// 844802 ns/op
func BenchmarkMemorySorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := genInt32Nums(1000, 1000*100)
		MemorySorter(nums)
	}
}

// 6438682600 ns/op
func BenchmarkHardSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := genInt32Nums(1000, 1000*100)
		HardSorter(nums)
	}
}

// 150286186 ns/op
func TestHardMergeSorterValid(t *testing.T) {
	nums := genInt32Nums(1000, 1000*100)
	sortedNums := HardMergeSorterValid(nums)
	if !sortedInt32(sortedNums) {
		t.Errorf("not sorted")
	}
}

func BenchmarkHardMergeSorter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums := genInt32Nums(1000, 1000*100)
		HardMergeSorter(nums)
	}
}
