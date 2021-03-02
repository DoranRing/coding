package algorithms

import "testing"

func Test_lastStoneWeightByHeap(t *testing.T) {
	type args struct {
		stones []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{stones: []int{2, 7, 4, 1, 8, 1}}, 1},
		{"2", args{stones: []int{8}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lastStoneWeightByHeap(tt.args.stones); got != tt.want {
				t.Errorf("lastStoneWeightByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
