package algorithms

import (
	"reflect"
	"testing"
)

func Test_reverseListByRecursion(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{&ListNode{1, &ListNode{2, &ListNode{3, nil}}}}, nil},
		{"1", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListByRecursion(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListByRecursion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseListByLoop(t *testing.T) {
	type args struct {
		head *ListNode
	}
	listNode := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{listNode}, listNode},
		{"1", args{nil}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListByLoop(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListByLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
