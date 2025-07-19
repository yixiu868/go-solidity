package main

import (
	"github.com/yixiu868/go-solidity/gobase/task2/lock"
	"time"
)

func main() {
	lock.AtomicQ2()

	time.Sleep(10 * time.Second)
}
