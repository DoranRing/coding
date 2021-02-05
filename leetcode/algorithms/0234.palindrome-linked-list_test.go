package algorithms

import "testing"

func Test_isPalindromeByCopy(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, false},
		{"2", args{&ListNode{1, &ListNode{2, &ListNode{1, nil}}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeByCopy(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeByCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeByRecursion(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, false},
		{"2", args{&ListNode{1, &ListNode{2, &ListNode{1, nil}}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeByRecursion(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeByRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindromeByReverse(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, false},
		{"2", args{&ListNode{1, &ListNode{2, &ListNode{1, nil}}}}, true},
		{"3", args{&ListNode{1, &ListNode{2, nil}}}, false},
		{"4", args{&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{1, nil}}}}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeByReverse(tt.args.head); got != tt.want {
				t.Errorf("isPalindromeByReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
