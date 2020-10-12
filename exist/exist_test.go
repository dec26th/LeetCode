package exist

import "testing"

func TestExist(t *testing.T) {
	test := [][]byte{
		{'a', 'b', 'c', 'e'},
		{'s', 'f', 'e', 's'},
		{'a', 'd', 'e', 'e'},
	}
	result := "abceseeefs"
	exist(test, result)
}
