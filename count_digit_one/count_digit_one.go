package count_digit_one

//输入一个整数 n ，求1～n这n个整数的十进制表示中1出现的次数。
//
//例如，输入12，1～12这些整数中包含1 的数字有1、10、11和12，1一共出现了5次。

func countDigitOne(n int) int {
	var result, oneToShow int
	layer := 1

	for n != 0 {
		left := n % 10
		n /= 10

		result += layer + oneToShow*left

		oneToShow = 10*oneToShow + layer
		layer *= 10
	}
	return result
}
