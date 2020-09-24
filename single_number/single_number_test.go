package single_number

import "testing"

func TestSingleNumber(t *testing.T) {
	test := []int{2, 2, 1}
	singleNumber(test)
}

func TestTwoNumberOnce(t *testing.T) {
	test := []int{1, 1, 2, 2, 3, 4}
	twoNumberOnce(test)
}
