package algorithms

import (
	"testing"
)

func Test_deleteDuplicatesByTwoPoint(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{head: &ListNode{1,
				&ListNode{1,
					&ListNode{2,
						&ListNode{3, nil}}}}},
			nil,
		},
		{
			"2",
			args{head: &ListNode{1,
				&ListNode{2,
					&ListNode{2,
						&ListNode{3, nil}}}}},
			nil,
		},
		{
			"2",
			args{head: &ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{3, nil}}}}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteDuplicatesByTwoPoint(tt.args.head)
		})
	}
}
