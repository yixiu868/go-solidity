package task1

import "sort"

// 合并区间
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	// 首先对区间按起始点排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged := [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		last := merged[len(merged)-1]

		// 如果当前区间的起始点小于等于前一个区间的结束点，则合并
		if interval[0] <= last[1] {
			if interval[1] > last[1] {
				merged[len(merged)-1][1] = interval[1]
			}
		} else {
			// 否则添加新区间
			merged = append(merged, interval)
		}
	}

	return merged
}
