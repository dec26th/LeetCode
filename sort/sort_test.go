package sort

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {
	test := []int{1, 4, 2, 6, 12, 1231241, 21, 2, 5, 3, 4}
	fmt.Println(QuickSort(test))
}


func TestHeapSort(t *testing.T) {
	test := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	buildMaxHeap(test, len(test))
	fmt.Println(test)
	fmt.Println(HeapSort(test))
}
