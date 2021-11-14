package heap

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	result := []int{9, 8,7,6,5,4,3,2,1,0}
	buildMinHeap(result)
	fmt.Println(result)
}

func TestMaxHeap(t *testing.T) {
	result := []int{1, 2, 3, 4, 5, 6, 7, 8}
	buildMaxHeap(result)
	fmt.Println(result)
}
