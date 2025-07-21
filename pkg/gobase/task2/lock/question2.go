package lock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int32
var wg sync.WaitGroup

func AtomicQ2() {
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				atomic.AddInt32(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("计数器", atomic.LoadInt32(&counter))
}
