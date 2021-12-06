package longest_palindrome

//给你一个字符串 s，找到 s 中最长的回文子串。

// out of time
func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= 0; j-- {
			if s[i] == s[j] && isLongestPalindrome(i + 1, j - 1, s){
				if j + 1 - i > len(result) {
					result = s[i:j+1]
				}
			}
		}
	}
	return result
}

func isLongestPalindrome(left, right int, s string) bool {
	if left >= right {
		return true
	}

	return s[left] == s[right] && isLongestPalindrome(left + 1, right - 1, s)
}

func longestPalindromeBetter(s string) string {
	left, right := 0, 0
	for i := 0; i < len(s); i++ {
		len1 := getPalindrome(s, i - 1, i + 1)
		len2 := getPalindrome(s, i, i+1)

		length := max(len1, len2)
		if length > right - left + 1 {
			if length % 2 == 0 {
				right = i + length / 2
				left = i - length / 2 + 1
			} else {
				right = i + length / 2
				left = i - length / 2
			}
		}
	}

	return s[left:right+1]
}

func max (x, y int) int {
	if x > y {
		return x
	}
	return y
}

func getPalindrome(s string, left, right int) int {
	for left < right && left >= 0 && right < len(s) && s[left] == s[right] {
		right++
		left--
	}

	return right - left - 1
}