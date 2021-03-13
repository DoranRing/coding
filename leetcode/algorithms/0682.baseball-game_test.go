package algorithms

import "testing"

func Test_calPointsByStack(t *testing.T) {
	type args struct {
		ops []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{[]string{"5", "2", "C", "D", "+"}}, 30},
		{"2", args{[]string{"5", "-2", "4", "C", "D", "9", "+", "+"}}, 27},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calPointsByStack(tt.args.ops); got != tt.want {
				t.Errorf("calPointsByStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
