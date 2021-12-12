package genarate_parenthesis

var result []string

func generateParenthesis(n int) []string {
	result = make([]string, 0)
	_generateParenthesis(n, 0, 0, "")
	return result
}

func _generateParenthesis(n, left, right int, parenthesis string) {
	if left == n && right == n {
		result = append(result, parenthesis)
		return
	}
	if left < n {
		appendLeft := parenthesis + "("
		_generateParenthesis(n, left+1, right, appendLeft)
	}

	if left > right {
		if right < n {
			appendRight := parenthesis + ")"
			_generateParenthesis(n, left, right+1, appendRight)
		}
	}
}

