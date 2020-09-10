package reverse_words

import "strings"

func reverseWords(s string) string {
	if s == "" {
		return ""
	}
	lenOfString := len(s)
	result := make([]string, 0)
	rightBound := lenOfString

	for nowBound := lenOfString - 1; nowBound >= 0; nowBound-- {

		if s[nowBound] == ' ' {
			if nowBound + 1 != rightBound {
				result = append(result, s[nowBound + 1: rightBound])
			}
			rightBound = nowBound
		}
	}

	if s[0] != ' ' {
		result = append(result, s[0 : rightBound])
	}

	return strings.Join(result, " ")
}

func reverseWordsWithAPI(s string) string {
	temp := strings.Split(s, " ")
	result := make([]string, 0)

	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] != "" {
			result = append(result, temp[i])
		}
	}

	return strings.Join(result, " ")
}
