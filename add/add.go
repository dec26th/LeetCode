package add

//写一个函数，求两个整数之和，要求在函数体内不得使用 “+”、“-”、“*”、“/” 四则运算符号。

func add(a int, b int) int {
	var result, offset int32
	carry := false
	offset = 1

	for i := 0; i < 32; i++ {
		if carry {
			result = result ^ (int32(a) & offset) ^ (int32(b) & offset) ^ offset
			carry = int32(a) & offset == offset || int32(b) & offset == offset
		} else {
			result = result ^ (int32(a) & offset) ^ (int32(b) & offset)
			carry = (int32(a) & offset & int32(b) & offset)  == offset
		}


		offset <<= 1
	}
	return int(result)
}
