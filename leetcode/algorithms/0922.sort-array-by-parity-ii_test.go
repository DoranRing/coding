package algorithms

import (
	"reflect"
	"testing"
)

func Test_sortArrayByParityIIByTwoPoint(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"1", args{A: []int{4, 2, 5, 7}}, []int{4, 5, 2, 7}},
		{"2", args{A: []int{4}}, []int{4}},
		{"3", args{A: []int{1, 2}}, []int{2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArrayByParityIIByTwoPoint(tt.args.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortArrayByParityIIByTwoPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
