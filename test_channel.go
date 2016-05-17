package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	go func() {
		fmt.Println("I am working")
		ch <- 1
		ch <- 2
		ch <- 3
		fmt.Println("Over")
		close(ch)
	}()

	for i := range ch {
		fmt.Println(i)
		fmt.Println(time.Now())
		time.Sleep(3 * time.Second)
	}

}
