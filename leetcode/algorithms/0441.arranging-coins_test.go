package algorithms

import "testing"

func Test_arrangeCoins(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5}, 2},
		{"2", args{8}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoinsByBinarySearch(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoinsByBinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrangeCoinsBySum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{5}, 2},
		{"2", args{8}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrangeCoinsBySum(tt.args.n); got != tt.want {
				t.Errorf("arrangeCoinsBySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
