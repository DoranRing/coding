package algorithms

import (
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{
				genListNode([]int{2, 4, 3}),
				genListNode([]int{5, 6, 4}),
			},
			genListNode([]int{7, 0, 8}),
		},
		{
			"2",
			args{
				genListNode([]int{0}),
				genListNode([]int{0}),
			},
			genListNode([]int{0}),
		},
		{
			"3",
			args{
				genListNode([]int{9, 9, 9, 9, 9, 9, 9}),
				genListNode([]int{9, 9, 9, 9}),
			},
			genListNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !listNodeEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
