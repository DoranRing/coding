package algorithms

import "testing"

func Test_hasCycleByHash(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, &ListNode{3, head}}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, false},
		{"2", args{head: head}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycleByHash(tt.args.head); got != tt.want {
				t.Errorf("hasCycleByHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasCycleByTwoPoint(t *testing.T) {
	head := &ListNode{1, nil}
	head.Next = &ListNode{2, &ListNode{3, head}}
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, false},
		{"2", args{head: head}, true},
		{"3", args{&ListNode{1, nil}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycleByTwoPoint(tt.args.head); got != tt.want {
				t.Errorf("hasCycleByTwoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
