package main

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	test := []int {1, 2, 3}
	copy(test[:2], test[1:])
	fmt.Println(test)
}
