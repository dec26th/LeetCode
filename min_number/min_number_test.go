package min_number

import (
	"fmt"
	"testing"
)

func TestMinNumber(t *testing.T) {
	temp := []int{
		20,
		1,
	}
	fmt.Println(minNumber(temp))
}

func TestLess(t *testing.T) {
	fmt.Println("100000000" < "3")
}
