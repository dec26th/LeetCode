package first_uniq_char

func firstUniqChar(s string) byte {
	charMap := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		if _, ok := charMap[s[i]]; !ok {
			charMap[s[i]] = 1
			continue
		}
		charMap[s[i]] += 1
	}

	for i := 0; i < len(s); i++ {
		if charMap[s[i]] == 1 {
			return s[i]
		}
	}

	return byte(' ')
}
