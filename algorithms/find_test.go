package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		sortNums []int
		val      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{}, 5}, -1},
		{"2", args{[]int{1, 2, 3, 4, 5, 6}, 5}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.sortNums, tt.args.val); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
