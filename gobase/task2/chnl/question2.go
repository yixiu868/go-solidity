package chnl

import "fmt"

func ChannelQ2() {
	ch := make(chan int, 5)

	go func() {
		for i := 1; i <= 100; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()
}
