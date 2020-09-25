package my_lru_cache

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	test := Constructor(2)
	test.Put(2, 1)
	test.Put(2, 2)
	fmt.Println(test.Get(2))
	test.Put(1, 1)
	test.Put(4, 1)
	fmt.Println(test.Get(2))
}
