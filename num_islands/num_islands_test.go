package num_islands

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T)  {
	test := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	fmt.Println(numIslands(test))
}
