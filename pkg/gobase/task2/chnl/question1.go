package chnl

import "fmt"

func ChannelQ1() {
	ch := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}
