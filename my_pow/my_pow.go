package my_pow

// 剑指 Offer 16. 数值的整数次方

// 实现函数double Power(double base, int exponent)，求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。


func myPow(x float64, n int) float64 { //timeout
	if x == 0 {
		return x
	}

	var result float64
	result = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}


