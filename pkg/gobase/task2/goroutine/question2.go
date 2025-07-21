package goroutine

import (
	"fmt"
	"reflect"
	"runtime"
	"sync"
	"time"
)

func ExecuteJobGroup() {
	slice := []func(){}
	slice = append(slice, hello, hello2, hello3)
	doExecuteJobGroup(slice)
}

func doExecuteJobGroup(sliceFunc []func()) {
	wg := sync.WaitGroup{}
	wg.Add(len(sliceFunc))

	for _, v := range sliceFunc {
		go func() {
			defer wg.Done()
			starttime := time.Now()
			v()
			executetime := time.Now().Sub(starttime)
			name := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
			fmt.Println("func", name, "execute time:", executetime)
		}()
	}

	wg.Wait()
}

func hello() {
	time.Sleep(1 * time.Second)
}

func hello2() {
	time.Sleep(10 * time.Second)
}

func hello3() {
	time.Sleep(5 * time.Second)
}
