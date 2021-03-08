package algorithms

import (
	"reflect"
	"testing"
)

func genData1() *ListNode {
	return &ListNode{1,
		&ListNode{2,
			&ListNode{3,
				&ListNode{4, nil}}}}
}

func Test_swapPairs(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{head: genData1()}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := swapPairs(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("swapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
