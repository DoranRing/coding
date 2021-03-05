package algorithms

import (
	"reflect"
	"testing"
)

func Test_allCellsDistOrder(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{R: 1, C: 2, r0: 0, c0: 0}, [][]int{{0, 0}, {0, 1}}},
		{"2", args{R: 2, C: 2, r0: 0, c0: 1}, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrder(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_allCellsDistOrderByBucket(t *testing.T) {
	type args struct {
		R  int
		C  int
		r0 int
		c0 int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{R: 1, C: 2, r0: 0, c0: 0}, [][]int{{0, 0}, {0, 1}}},
		{"2", args{R: 2, C: 2, r0: 0, c0: 1}, [][]int{{0, 1}, {0, 0}, {1, 1}, {1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allCellsDistOrderByBucket(tt.args.R, tt.args.C, tt.args.r0, tt.args.c0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allCellsDistOrderByBucket() = %v, want %v", got, tt.want)
			}
		})
	}
}
