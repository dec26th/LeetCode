package merge

import "testing"

func TestMerge(t *testing.T) {
	test := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	merge(test)
}
