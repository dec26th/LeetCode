package max_profit

import "testing"

func Test_maxProfitK(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
		{
			name: "pass",
			args: args{
				k:      2,
				prices: []int{3, 3, 5, 0, 0, 3, 1, 4},
			},
			want: 6,
		},
		{
			name: "pass",
			args: args{
				k:      2,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0},
			},
			want: 13,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitK(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfitK() = %v, want %v", got, tt.want)
			}
		})
	}
}
