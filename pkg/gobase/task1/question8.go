package task1

// 两数之和
func TwoSum(nums []int, target int) []int {
	maps := make(map[int]int)

	for i, v := range nums {
		maps[v] = i
	}

	for i, v := range nums {
		if maps[target-v] != 0 && maps[target-v] != i {
			return []int{i, maps[target-v]}
		}
	}
	return []int{}
}
