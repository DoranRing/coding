package algorithms

import "testing"

func Test_maxDepthByAfterOrder(t *testing.T) {
	type args struct {
		node *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{&TreeNode{10, &TreeNode{5, &TreeNode{1, nil, nil}, nil}, nil}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.node); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
