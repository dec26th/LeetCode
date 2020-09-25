package my_lru_cache

import "testing"

func TestLRUCache(t *testing.T) {
	test := Constructor(10)
	test.Head.Next = &Node{
		Val:  1,
		Key:  2,
		Pre:  test.Head,
		Next: nil,
	}

}
