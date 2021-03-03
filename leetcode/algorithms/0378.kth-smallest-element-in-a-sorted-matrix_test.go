package algorithms

import "testing"

func Test_kthSmallestByHeap(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{
				matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}},
				k:      8,
			},
			13,
		},
		{
			"2",
			args{
				matrix: [][]int{{-5}},
				k:      1,
			},
			-5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallestByHeap(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallestByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
