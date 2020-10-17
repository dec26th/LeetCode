package hamming_weight

// 剑指 Offer 15. 二进制中1的个数

//请实现一个函数，输入一个整数，输出该数二进制表示中 1 的个数。
//例如，把 9表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。

func hammingWeight(num uint32) int {
	result := 0
	for i := 0; i < 32; i++ {
		if num&1 == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}

func hammingWeightBetter(num uint32) int {
	result := 0

	for num != 0 {
		result++
		num &= num - 1
	}
	return result
}
