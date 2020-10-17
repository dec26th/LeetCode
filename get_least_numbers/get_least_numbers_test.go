package get_least_numbers

import (
	"fmt"
	"testing"
)

func TestGetLeastNumbers(t *testing.T) {
	test := []int {2, 0, 1, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(getLeastNumbers(test, 2))
}
