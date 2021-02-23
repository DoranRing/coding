package algorithms

import "testing"

func Test_maxTurbulenceSizeBySlidingWindow(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []int{9, 4, 2, 10, 7, 8, 8, 1, 9}}, 5},
		{"2", args{arr: []int{4, 8, 12, 16}}, 2},
		{"3", args{arr: []int{100}}, 1},
		{"4", args{arr: []int{0, 1, 1, 0, 1, 0, 1, 1, 0, 0}}, 5},
		{"5", args{arr: []int{0, 8, 45, 88, 48, 68, 28, 55, 17, 24}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxTurbulenceSizeBySlidingWindow(tt.args.arr); got != tt.want {
				t.Errorf("maxTurbulenceSizeBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
