package search

import "testing"

func Test_searchToday(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 0},
			want: 4,
		},
		{
			name: "pass",
			args: args{nums: []int{4, 5, 6, 7, 0, 1, 2}, target: 3},
			want: -1,
		},
		{
			name: "pass",
			args: args{nums: []int{1}, target: 0},
			want: -1,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchToday(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchToday() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchToday1(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "pass",
			args: args{
				nums:   []int{1, 3},
				target: 2,
			},
			want: -1,
		},
		{
			name: "pass",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "pass",
			args: args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: 4,
		},
		{
			name: "pass",
			args: args{
				nums:   []int{1, 3},
				target: 3,
			},
			want: 1,
		},
		{
			name: "pass",
			args: args{
				nums:   []int{3, 1},
				target: 1,
			},
			want: 1,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchToday(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchToday() = %v, want %v", got, tt.want)
			}
		})
	}
}
