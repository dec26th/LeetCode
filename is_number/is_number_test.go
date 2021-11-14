package is_number

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestRegChars(t *testing.T) {
	test := "123asd123asd"
	test2 := "1231231e12312312"

	fmt.Println(regChars.MatchString(test))
	fmt.Println(regChars.MatchString(test2))
}

func TestStrconv(t *testing.T) {
	test := "-1123"
	result, err := strconv.Atoi(test)
	fmt.Println(result, err)
}

func TestSplit(t *testing.T) {
	s := ".1"

	result := strings.Split(s, ".")
	for _, v := range result {
		fmt.Println(v)
	}
}
