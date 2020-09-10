package check_inclusion

//todo 排列问题
func checkInclusion(s1 string, s2 string) bool  {
	lenOfS1 := len(s1)
	lenOfS2 := len(s2)
	if lenOfS1 > lenOfS2 {
		return false
	}

	charSet1, charSet2 := [26]int{}, [26]int{}

	for i := 0; i < lenOfS1; i++ {
		charSet1[int(s1[i] - 'a')] ++
		charSet2[int(s2[i] - 'a')] ++
	}

	for j := 0; j < lenOfS2 - lenOfS1; j++ {
		if charSet1 == charSet2 {
			return true
		}
		charSet2[s2[j] - 'a']--
		charSet2[s2[j + lenOfS1] - 'a'] ++
	}

	return charSet1 == charSet2
}

