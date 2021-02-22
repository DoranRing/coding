package algorithms

import "testing"

func Test_checkInclusionBySlidingWindow(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{s1: "ab", s2: "eidbaooo"}, true},
		{"1", args{s1: "aab", s2: "eidboaoo"}, false},
		{"1", args{s1: "horse", s2: "ros"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkInclusionBySlidingWindow(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("checkInclusionBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
