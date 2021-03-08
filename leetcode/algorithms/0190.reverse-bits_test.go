package algorithms

import (
	"testing"
)

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"1", args{0b00000010100101000001111010011100}, 0b00111001011110000010100101000000},
		{"2", args{0b11111111111111111111111111111101}, 0b10111111111111111111111111111111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBitsByMath(tt.args.num); got != tt.want {
				t.Errorf("reverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBitsByBit(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{"1", args{0b00000010100101000001111010011100}, 0b00111001011110000010100101000000},
		{"2", args{0b11111111111111111111111111111101}, 0b10111111111111111111111111111111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBitsByBit(tt.args.num); got != tt.want {
				t.Errorf("reverseBitsByBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_reverseBitsByBit(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
