package main

import (
	"fmt"
	"time"
)

func main() {
	channelWithCache()
	channelWithOutCache()
}

func channelWithCache() {
	ch := make(chan string, 1)

	go func() {
		//ch 方向 从右往做看，比较顺眼
		ch <- "hello,first msg into channel"
		time.Sleep(1 * time.Second)
		ch <- "hello,second msg into channel"
	}()
	time.Sleep(2 * time.Second)
	msg := <-ch
	fmt.Println(time.Now().String() + msg)
	msg = <-ch
	fmt.Println(time.Now().String() + msg)
}

func channelWithOutCache() {
	ch := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		ch <- "hello,first msg into out cache channel"
	}()

	msg := <-ch
	fmt.Println(msg)
}
