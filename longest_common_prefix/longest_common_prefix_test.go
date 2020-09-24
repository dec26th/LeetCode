package longest_common_prefix

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	test := "hello"
	for k, v := range test {
		fmt.Printf("第%d 个的字符为 %c", k, v)
	}
}

func TestLongestCommonPrefix(t *testing.T) {
	test := []string{"aa", "ab"}
	longestCommonPrefix(test)
}
