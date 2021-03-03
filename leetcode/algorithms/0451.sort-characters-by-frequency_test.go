package algorithms

import "testing"

func Test_frequencySortByHeap(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{s: "tree"}, "eert"},
		{"2", args{s: "cccaaa"}, "cccaaa"},
		{"3", args{s: "Aabb"}, "bbAa"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySortByHeap(tt.args.s); got != tt.want {
				t.Errorf("frequencySortByHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
