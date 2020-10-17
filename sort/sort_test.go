package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	test := []int{1, 4, 2, 6, 12, 1231241, 21, 2, 5, 3, 4}
	fmt.Println(QuickSort(test))
}
