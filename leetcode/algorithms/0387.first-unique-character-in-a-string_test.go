package algorithms

import "testing"

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{"1", args{"leetcode"}, 0},
}

func Test_firstUniqChar(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqCharByArray(tt.args.s); got != tt.want {
				t.Errorf("firstUniqChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstUniqCharByHash(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstUniqCharByHash(tt.args.s); got != tt.want {
				t.Errorf("firstUniqCharByHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
