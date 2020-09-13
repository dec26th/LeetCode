package lru_cache

import (
	"fmt"
	"testing"
)

func TestLruCache(t *testing.T) {
	test := Constructor(2)
	test.Put(2, 1)
	test.Put(1, 1)
	test.Put(2, 3)
	test.Put(4, 1)
	fmt.Println(test.Get(1))
	fmt.Println(test.Get(2))
}


func TestCopy(t *testing.T) {
	test := []int {1, 2, 3, 4, 5, 6}
	copy(test[0:], test[1:])
	fmt.Println(test)
}