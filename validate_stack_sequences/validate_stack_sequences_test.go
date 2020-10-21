package validate_stack_sequences

import (
	"fmt"
	"testing"
)

func TestValidateStackSequences(t *testing.T) {
	push := []int {1, 2, 3, 4, 5}
	popped := []int {4, 5, 3, 2, 1}

	fmt.Println(validateStackSequences(push, popped))
}
