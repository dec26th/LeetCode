package valid_UTF8


func validUtf8(data []int) bool {
	if len(data) == 0 {
		return false
	}

	byteLength := getByteLength(data[0])
	if byteLength == -1 {
		return false
	}

	for i := 1; i < len(data); i++ {

		if data[i] & (3 << 6) != 1 << 7 {
			continue
		}
		byteLength--
	}
	return byteLength == 0
}

func getByteLength(data int) int {
	if data >> 7 == 0 {
		return 0
	}

	if data >> 5 == 6 {
		return 1
	}

	if data >> 4 == 14 {
		return 2
	}

	if data >> 3 == 30 {
		return 3
	}

	return -1
}

