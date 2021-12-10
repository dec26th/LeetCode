package sort

import (
	"reflect"
	"testing"
)

func Test_myQuickSort(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "pass",
			args: args{nums: []int{1, 346, 345, 34567, 34, 7, 312, 14, 4}},
			want: []int{1, 4, 7, 14, 34, 312, 345, 346, 34567},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myQuickSort(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("myQuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
