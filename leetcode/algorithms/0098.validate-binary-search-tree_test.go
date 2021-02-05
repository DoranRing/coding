package algorithms

import "testing"

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&TreeNode{0, nil, nil}}, true},
		{"2", args{&TreeNode{-2147483648, nil, nil}}, true},
		{
			"3",
			args{
				&TreeNode{5,
					&TreeNode{1, nil, nil},
					&TreeNode{4,
						&TreeNode{3, nil, nil},
						&TreeNode{6, nil, nil},
					},
				},
			},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
