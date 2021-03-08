package algorithms

import "testing"

func Test_isPalindromeByTwoPoint(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"A man, a plan, a canal: Panama"}, true},
		{"2", args{"race a car"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeByTwoPoint(tt.args.s); got != tt.want {
				t.Errorf("isPalindromeByTwoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
