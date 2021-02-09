package algorithms

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			"1",
			args{[]*ListNode{
				genListNode([]int{1, 4, 5}),
				genListNode([]int{1, 3, 4}),
				genListNode([]int{2, 6}),
			}},
			genListNode([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{"2", args{[]*ListNode{}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
