package remove_nth_from_end

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "pass",
			args: args{head: &ListNode{Val: 1}, n: 1},
			want: nil,
		},
		{
			name: "test",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, n: 2},
			want: &ListNode{Val: 2},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
