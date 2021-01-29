package algorithms

import "testing"

func Test_missingNumberByMath(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{3, 0, 1}}, 2},
		{"2", args{[]int{0, 1}}, 2},
		{"3", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumberByMath(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberByMath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_missingNumberByXOR(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"1", args{[]int{3, 0, 1}}, 2},
		{"2", args{[]int{0, 1}}, 2},
		{"3", args{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}}, 8},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := missingNumberByXOR(tt.args.nums); got != tt.want {
				t.Errorf("missingNumberByXOR() = %v, want %v", got, tt.want)
			}
		})
	}
}
