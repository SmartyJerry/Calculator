package calc

import "testing"

func TestAddition_Calculate(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Zeros", args{0, 0}, 0},
		{"Positives", args{1, 5}, 6},
		{"Negative First", args{-2, 5}, 3},
		{"Negative Second", args{2, -5}, -3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addition := Addition{}
			if got := addition.Calculate(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
