package find_number_in_2darray

import "testing"

func TestFindNUmber(t *testing.T) {
	Array1 := make([]int, 1)
	Array1[0] = 10
	Array := make([][]int, 1)
	Array[0] = Array1
	print(findNumberIn2DArrayLine(Array, -10))
}
