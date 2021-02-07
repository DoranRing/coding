package algorithms

import (
	"reflect"
	"testing"
)

func Test_levelOrderByQueue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"1",
			args{
				&TreeNode{
					10,
					&TreeNode{5, nil, nil},
					&TreeNode{
						15,
						&TreeNode{
							25,
							&TreeNode{18, nil, nil},
							nil,
						},
						nil,
					},
				},
			},
			[][]int{{10}, {5, 15}, {25}, {18}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderByQueue(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
