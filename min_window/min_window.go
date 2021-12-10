package min_window

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
//https://leetcode-cn.com/problems/minimum-window-substring/

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}
	mapping := make(map[uint8]int, len(t))
	for i := 0; i < len(t); i++ {
		mapping[t[i]]++
	}

	record := len(t)
	left, right := 0, len(s)
	exist := false

	i, j := 0, 0
	for i < len(s) && j < len(s) {
		if num, ok := mapping[s[j]]; ok {
			if num > 0 {
				record--
			}
			mapping[s[j]]--

		}

		if record == 0 {
			for ; i <= j; i++ {

				if num, ok := mapping[s[i]]; ok {

					if num == 0 {
						if j-i < right-left {
							right = j
							left = i
							exist = true
						}
						mapping[s[i]]++
						break
					}
					mapping[s[i]]++
				}
			}

			record++
			i++
		}
		j++
	}

	if !exist {
		return ""
	}

	return s[left : right+1]
}
