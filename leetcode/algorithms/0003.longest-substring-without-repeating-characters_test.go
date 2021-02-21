package algorithms

import "testing"

func Test_lengthOfLongestSubstringBySlidingWindow(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{s: "abcabcbb"}, 3},
		{"2", args{s: "bbbbb"}, 1},
		{"3", args{s: "pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringBySlidingWindow(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringBySlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
