package algorithms

import "testing"

func Test_reverseVowels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{s: "hello"}, "holle"},
		{"2", args{s: "leetcode"}, "leotcede"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowelsByCollisionPointer(tt.args.s); got != tt.want {
				t.Errorf("reverseVowelsByCollisionPointer() = %v, want %v", got, tt.want)
			}
		})
	}
}
