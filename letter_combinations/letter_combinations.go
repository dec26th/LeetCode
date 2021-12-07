package letter_combinations

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	var numToLetter = map[uint8][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	if len(digits) == 1 {
		return numToLetter[digits[0]]
	}

	temp := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		cur := numToLetter[digits[i]]

		if len(temp) == 0 {
			temp = cur
		} else {
			tt := make([]string, 0, len(temp)*len(cur))
			for j := 0; j < len(temp); j++ {
				for k := 0; k < len(cur); k++ {
					tt = append(tt, fmt.Sprintf("%s%s", temp[j], cur[k]))
				}
			}

			temp = tt
		}
	}

	return temp
}
