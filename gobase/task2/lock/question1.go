package lock

import (
	"fmt"
	"sync"
)

func SyncQ1() {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(10)

	counter := 0

	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 1000; i++ {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("计数器", counter)
}
