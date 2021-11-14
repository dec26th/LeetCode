package cutting_rope

// 剑指 Offer 14- I. 剪绳子

// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。
// 请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}

	a, b := n/3, n%3

	switch b {
	case 0:
		return pow(3, a) % 1000000007
	case 1:
		return (pow(3, a-1) * 4) % 1000000007
	case 2:
		return (pow(3, a) * 2) % 1000000007
	}
	return 0
}

func pow(a, b int) int {
	result := 1
	for i := 0; i < b; i++ {
		result = (result * a) % 1000000007
	}
	return result
}
