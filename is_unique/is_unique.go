package is_unique


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
	var mark int
	lenOfAstr := len(astr)
	asciiA := int('a')

	for i := 0; i < lenOfAstr; i++ {
		index := int(rune(astr[i])) - asciiA
		if mark & (1 << index) != 0 {
			return false
		}

		mark |= 1 << index
	}
	return true
}