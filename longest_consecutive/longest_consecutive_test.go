package longest_consecutive

import (
	"fmt"
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	test := []int{-1, 0, 1}
	fmt.Println(longestConsecutive(test))
}

func Test_longestConsecutive(t *testing.T) {
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
			args: args{nums: []int{0, 0, -1}},
			want: 2,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestConsecutive(tt.args.nums); got != tt.want {
				t.Errorf("longestConsecutive() = %v, want %v", got, tt.want)
			}
		})
	}
}
