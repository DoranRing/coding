package algorithms

import "testing"

func Test_maxAreaByForce(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"2", args{[]int{1, 1}}, 1},
		{"3", args{[]int{4, 3, 2, 1, 4}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaByForce(tt.args.height); got != tt.want {
				t.Errorf("maxAreaByForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxAreaByCollisionPointer(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"2", args{[]int{1, 1}}, 1},
		{"3", args{[]int{4, 3, 2, 1, 4}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaByCollisionPointer(tt.args.height); got != tt.want {
				t.Errorf("maxAreaByCollisionPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
