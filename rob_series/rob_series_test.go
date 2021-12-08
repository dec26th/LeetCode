package rob_series

import "testing"

func Test_rob20211208(t *testing.T) {
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
			args: args{nums: []int{2,1,1,2}},
			want: 4,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob20211208(tt.args.nums); got != tt.want {
				t.Errorf("rob20211208() = %v, want %v", got, tt.want)
			}
		})
	}
}
