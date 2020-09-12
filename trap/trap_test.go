package trap

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	test := []int{2, 1, 0, 4}
	fmt.Println(trap(test))
}

func TestMap(t *testing.T) {
	test := make(map[string]string)

	test["a"] = "a"
	Trick(test)
	fmt.Println(test)
}

func Trick(test map[string]string) {
	test["a"] = "b"
}