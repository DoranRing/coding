package algorithms

import "testing"

func Test_isAnagramBySort(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"anagram", "nagaram"}, true},
		{"2", args{"rat", "car"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagramBySort(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramBySort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagramByHash(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"anagram", "nagaram"}, true},
		{"2", args{"rat", "car"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagramByHash(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramByHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isAnagramByArray(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"anagram", "nagaram"}, true},
		{"2", args{"rat", "car"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagramByArray(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isAnagramByArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
