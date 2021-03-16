package algorithms

import (
	"reflect"
	"testing"
)

func Test_countBits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{num: 2}, []int{0, 1, 1}},
		{"2", args{num: 5}, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countBitsByCompute(tt.args.num); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countBitsByCompute() = %v, want %v", got, tt.want)
			}
		})
	}
}
