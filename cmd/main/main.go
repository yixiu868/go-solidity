package main

import "github.com/yixiu868/go-solidity/gobase/task1"

func main() {
	for i := range task1.Merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}) {
		print(i)
	}
}
