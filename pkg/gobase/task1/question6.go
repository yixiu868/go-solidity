package task1

// 删除有序数组中重复项
func RemoveDuplicates(nums []int) int {
	maps := make(map[int]int)
	length := len(nums)
	for i := 0; i < length; i++ {
		if maps[nums[i]] == 0 {
			maps[nums[i]] = 1
		} else {
			nums = append(nums[:i], nums[i+1:]...)
			i--
			length--
		}
	}
	return length
}
