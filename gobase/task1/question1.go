package task1

// 只出现一次的数字
func Question1(list []int) int {
	maps := make(map[int]int)
	for _, v := range list {
		maps[v]++
	}
	for k, v := range maps {
		if v == 1 {
			return k
		}
	}
	return -1
}
