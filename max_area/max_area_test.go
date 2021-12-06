package max_area

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass",
			args: args{height: []int{1, 1}},
			want: 1,
		},
		{
			name: "pass",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "pass",
			args: args{height: []int{4, 3, 2, 1, 4}},
			want: 16,
		},
		{
			name: "pass",
			args: args{height: []int{1, 2, 1}},
			want: 2,
		},
		{
			name: "pass",
			args: args{height: []int{1,8,6,2,5,4,8,25,7}},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxAreaBetter(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "pass",
			args: args{height: []int{1, 1}},
			want: 1,
		},
		{
			name: "pass",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "pass",
			args: args{height: []int{4, 3, 2, 1, 4}},
			want: 16,
		},
		{
			name: "pass",
			args: args{height: []int{1, 2, 1}},
			want: 2,
		},
		{
			name: "pass",
			args: args{height: []int{1,8,6,2,5,4,8,25,7}},
			want: 49,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxAreaBetter(tt.args.height); got != tt.want {
				t.Errorf("maxAreaBetter() = %v, want %v", got, tt.want)
			}
		})
	}
}