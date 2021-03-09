package algorithms

import (
	"reflect"
	"testing"
)

func Test_rotateRightByLoop(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1",
			args{&ListNode{1,
				&ListNode{2,
					&ListNode{3,
						&ListNode{4,
							&ListNode{5, nil},
						},
					},
				},
			}, 11,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateRightByLoop(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateRightByLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
