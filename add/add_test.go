package add

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	fmt.Println(add(1, -2))
}


func TestDouble(t *testing.T) {
	test := [][]int{{1, 2, 3, 4}, {11, 12, 13, 14}, {21, 22, 23, 24}, {31, 32, 33, 34}}
	test1 := test[:1]
	fmt.Println(test1[0:3])
	fmt.Println(test[:1][0:3])
}