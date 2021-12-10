package min_window

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pass",
			args: args{
				s: "adobecodebanc",
				t: "abc",
			},
			want: "banc",
		},
		{
			name: "pass",
			args: args{
				s: "a",
				t: "b",
			},
			want: "",
		},
		{
			name: "pass",
			args: args{
				s: "ab",
				t: "b",
			},
			want: "b",
		},
		{
			name: "pass",
			args: args{
				s: "ab",
				t: "a",
			},
			want: "a",
		},
		{
			name: "pass",
			args: args{
				s: "aa",
				t: "aa",
			},
			want: "aa",
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
