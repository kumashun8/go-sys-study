package main

import (
	"fmt"
	"sync"
)

func initialize() {
	fmt.Println("initialize")
}

var once sync.Once

func main() {
	once.Do(initialize)
	// 2回目以降は実行されない
	once.Do(initialize)
	once.Do(initialize)
}
