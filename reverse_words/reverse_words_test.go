package reverse_words

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T)  {
	test := "  hello world!  "
	fmt.Println(reverseWords(test))
}

func TestReverseWordsSecond(t *testing.T) {
	test := "  hello world!   "
	fmt.Println(reverseWordsWithAPI(test))
}