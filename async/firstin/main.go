package main

import (
	"fmt"
	"time"
)

var skipCnt int

func fanIn(ch1, ch2 <-chan string) <-chan string {
	newCh := make(chan string)
	go func() {
		for {
			select {
			case s := <-ch1:
				newCh <- s
			case s := <-ch2:
				newCh <- s
			default:
				skipCnt++
			}
		}
	}()
	return newCh
}

func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s:%d", msg, i)
			time.Sleep(time.Second)
		}
	}()
	return ch
}

func main() {
	ch := fanIn(generator("Hello"), generator("Bye"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}
	fmt.Println("Skipped:", skipCnt)
}
