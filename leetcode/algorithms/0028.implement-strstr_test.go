package algorithms

import "testing"

func Test_strStrByTwoPoint(t *testing.T) {
	type args struct {
		haystack string
		needle   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{"hello", "ll"}, 2},
		{"2", args{"aaaaa", "bba"}, -1},
		{"3", args{"", ""}, 0},
		{"4", args{"a", "a"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStrByTwoPoint(tt.args.haystack, tt.args.needle); got != tt.want {
				t.Errorf("strStrByTwoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
