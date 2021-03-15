package algorithms

import (
	"reflect"
	"testing"
)

func Test_copyRandomListByHash(t *testing.T) {
	node4 := &Node{1, nil, nil}
	node3 := &Node{10, node4, nil}
	node2 := &Node{11, node3, nil}
	node1 := &Node{13, node2, nil}
	node0 := &Node{7, node1, nil}
	node1.Random = node0
	node2.Random = node4
	node3.Random = node2
	node4.Random = node0
	type args struct {
		head *Node
	}
	tests := []struct {
		name string
		args args
		want *Node
	}{
		{
			"1",
			args{
				head: node0,
			},
			nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyRandomListByHash(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyRandomListByHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
