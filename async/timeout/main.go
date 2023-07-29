package main

import (
	"fmt"
	"time"
)

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s:%d", msg, i)
			time.Sleep(999 * time.Millisecond)
		}
	}()
	return ch
}

func main() {
	ch := generator("Hello")
	for i := 0; i < 10; i++ {
		select {
		case s := <-ch:
			fmt.Println(s)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}
