package maxProduct

import "testing"

func Test_maxProduct(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass",
			args: args{nums: []int{2, 3, -2, 4}},
			want: 6,
		},
		{
			name: "pass",
			args: args{nums: []int{-2, 0, -1}},
			want: 0,
		},
		{
			name: "pass",
			args: args{nums: []int{-3, -1, -1}},
			want: 3,
		},
		{
			name: "pass",
			args: args{nums: []int{0, 2}},
			want: 2,
		},
		{
			name: "pass",
			args: args{nums: []int{3, -1, 4}},
			want: 4,
		},
		{
			name: "pass",
			args: args{nums: []int{-4, -3, -2}},
			want: 12,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
