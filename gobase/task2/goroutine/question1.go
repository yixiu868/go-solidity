package goroutine

import (
	"fmt"
	"sync"
)

func ParallelPrintNum() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println("goroutine1:", i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println("goroutine2:", i)
		}
	}()

	wg.Wait()
}
