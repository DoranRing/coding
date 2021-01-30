package algorithms

import (
	"reflect"
	"testing"
)

func Test_generateByLoop(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{1}, [][]int{{1}}},
		{"2", args{2}, [][]int{{1}, {1, 1}}},
		{"3", args{3}, [][]int{{1}, {1, 1}, {1, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateByLoop(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateByLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
