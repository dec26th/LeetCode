package validate_stack_sequences

func validateStackSequences(pushed []int, popped []int) bool {
	num := 0
	index := 0
	lenOfPushed := len(pushed)
	mark := make(map[int]byte)

	for i := 0; i < lenOfPushed; i++ {
		if _, ok := mark[i]; !ok && pushed[i] == popped[index] {
			mark[i] = '1'
			index ++
			for ; i >= 0; i-- {
				if _, ok := mark[i]; !ok {
					i -= 1
					break
				}
			}
			num += 1
			continue
		}

	}

	return num == len(pushed)
}

//todo optimize