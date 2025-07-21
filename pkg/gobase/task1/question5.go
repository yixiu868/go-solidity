package task1

// 加1
func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}
	return digits
}
