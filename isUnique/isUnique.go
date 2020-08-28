package isUnique


func isUniqueMap(astr string) bool {
	lenOfAstr := len(astr)
	astrMap := make(map[string]int, lenOfAstr)

	for i := 0; i < lenOfAstr; i++ {
		if _, ok := astrMap[string(astr[i])]; ok == true {
			return false
		}
		astrMap[string(astr[i])] = i
	}
	return true
}

func isUniqueBitArray(astr string) bool {
	var mark int8
	lenOfAstr := len(astr)
	asciiA := int8('a')

	for i := 0; i < lenOfAstr; i++ {
		index := int8(rune(astr[i])) - asciiA
		if mark & (1 << index) != 0 {
			return false
		}

		mark |= 1 << index - asciiA
	}
	return true
}