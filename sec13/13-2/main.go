package main

import (
	"fmt"
	"sync"
)

func sub1(c int, wg *sync.WaitGroup) {
	fmt.Println("share by argument:", c)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	// 引数渡し
	go sub1(10, &wg)
	// クロージャのキャプチャ渡し
	c := 20
	wg.Add(1)
	go func() {
		fmt.Println("share by capture:", c*c)
		wg.Done()
	}()
	wg.Wait()

	tasks := []string{
		"cmake ..",
		"cmake . --build Release",
		"cpack",
	}
	for _, task := range tasks {
		wg.Add(1)
		go func() {
			// taskの値をキャプチャしているため、
			// goroutineが実行されるまでにtaskの値が変わってしまう
			fmt.Println(task)
			wg.Done()
		}()
	}
	wg.Wait()
}
