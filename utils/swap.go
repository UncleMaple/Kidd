package utils

// 传入两个数，传出小数和大数
func Swap(a, b int) (min, max int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	return min, max
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}
