package algorithms

import "testing"

func Test_isValidByStack(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"()[]{}"}, true},
		{"2", args{"(]"}, false},
		{"3", args{"([)]"}, false},
		{"4", args{"{[]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidByStack(tt.args.s); got != tt.want {
				t.Errorf("isValidByStack() = %v, want %v", got, tt.want)
			}
		})
	}
}
