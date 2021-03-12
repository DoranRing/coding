package algorithms

import (
	"testing"
)

func Test_reverseBetweenByHeadFirst(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{head: GenListNode([]int{1, 2, 3, 4, 5}), left: 2, right: 4},
			GenListNode([]int{1, 4, 3, 2, 5}),
		},
		{
			"2",
			args{head: GenListNode([]int{1, 2, 3, 4, 5}), left: 1, right: 3},
			GenListNode([]int{3, 2, 1, 4, 5}),
		},
		{
			"5",
			args{head: GenListNode([]int{1, 2, 3, 4, 5}), left: 1, right: 1},
			GenListNode([]int{1, 2, 3, 4, 5}),
		},
		{
			"6",
			args{head: GenListNode([]int{1, 2, 3, 4, 5}), left: 3, right: 5},
			GenListNode([]int{1, 2, 5, 4, 3}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetweenByHeadFirst(tt.args.head, tt.args.left, tt.args.right); !ListNodeEqual(got, tt.want) {
				t.Errorf("reverseBetweenByHeadFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
