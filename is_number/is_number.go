package is_number

import (
	"regexp"
	"strconv"
	"strings"
)

var regChars = regexp.MustCompile("[*a-df-zA-DF-Z]")

func isNumber(s string) bool {
	if isInValid(s) {
		return false
	}

	s = strings.Join(strings.Split(s, " "), "")
	if isInterger(s) || isPoint(s) {
		return true
	}
	return false
}

// 若字符串当中包含有多个小数点
func isInValid(s string) bool {
	return strings.Count(s, ".") > 1 || strings.Count(s, "e") > 1 || strings.Count(s, "E") > 1 ||
		regChars.MatchString(s)
}

func isInterger(s string) bool {
	if strings.Contains(s, ".") {
		return false
	}

	result := strings.Split(s, "e")
	if len(result) == 2 { // 包含e
		if _, err1 := strconv.Atoi(result[0]); err1 != nil {
			return false
		}

		if _, err2 := strconv.Atoi(result[1]); err2 != nil {
			return false
		}
		return true
	} else if len(result) == 1 { // 不包含e
		result := strings.Split(s, "E")

		if len(result) == 1 { // 也不包含E
			if _, err := strconv.Atoi(s); err != nil {
				return false
			}
			return true
		}

		if len(result) == 2 { // 包含E
			if _, err1 := strconv.Atoi(result[0]); err1 != nil {
				return false
			}

			if _, err2 := strconv.Atoi(result[1]); err2 != nil {
				return false
			}
			return true
		}
		return false

	}
	return false
}

func isPoint(s string) bool {
	if strings.Contains(s, ".") {
		result := strings.Split(s, ".")
		if len(result) > 2 {
			return false
		}
		if result[0] == "" {
			if _, err1 := strconv.Atoi(result[1]); err1 != nil {
				return false
			}
			return true
		}

		if result[1] == "" {
			if _, err1 := strconv.Atoi(result[0]); err1 != nil {
				return false
			}
			return true
		}
		if _, err1 := strconv.Atoi(result[0]); err1 != nil {
			return false
		}

		if _, err2 := strconv.Atoi(result[1]); err2 != nil {
			return false
		}
		return true
	}
	return false
}
